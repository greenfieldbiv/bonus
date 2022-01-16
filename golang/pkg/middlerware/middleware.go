package middlerware

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/service"
	"bivbonus/pkg/security"
	"bivbonus/pkg/security/jwt"
	"bivbonus/pkg/security/jwt/storage"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

const (
	contentTypeJson = "application/json"
	accessTokenKey  = "a_tkn"
	refreshTokenKey = "r_tkn"
)

type IMiddleWare interface {
	wareAll(ctx *fasthttp.RequestCtx)
	WareLogin(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler
	WareLogout(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler
	WareRegistry(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler
	WareCommon(next fasthttp.RequestHandler) fasthttp.RequestHandler
	WareSecurity(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler
}

type authMiddleWare struct {
	IMiddleWare
	storage      *storage.JwtStorage
	service      *service.Service
	passwordHash *security.PasswordHash
}

func NewMiddleWare(storage *storage.JwtStorage, service *service.Service) *authMiddleWare {
	return &authMiddleWare{storage: storage, passwordHash: &security.PasswordHash{}, service: service}
}

func (m authMiddleWare) WareSecurity(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer m.wareAll(ctx)
		request := &ctx.Request
		extractToken := security.ExtractToken(request)
		if extractToken == "" {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}
		// Если токен валидный, то делаем, что делали
		jwtStorage := security.JwtStorage()
		if accessTokenMetadataExtracted, err := security.ExtractAccessTokenMetadataExtracted(request, false); err == nil {
			// в хранилище с токенами он тоже должен быть (например, после логаута он пропадет)
			if jwtStorage.IsExistAccessToken(accessTokenMetadataExtracted.AccountId, accessTokenMetadataExtracted.AccessToken.Raw) {
				refreshTokenStr := jwtStorage.GetRefreshTokenById(strconv.FormatInt(accessTokenMetadataExtracted.AccountId, 10))
				if refreshToken, err := security.GetRefreshTokenByStr(refreshTokenStr, true); err == nil && security.TokenValid(refreshToken) {
					ctx.SetUserValue(security.ContextKey().AccountIdKey, accessTokenMetadataExtracted.AccountId)
					next(ctx)
					return
				} else {
					go jwtStorage.DeleteRefreshToken(accessTokenMetadataExtracted.AccountId, refreshTokenStr)
				}
			}
		} else {
			accessTokenMetadata, _ := security.ExtractAccessTokenMetadata(extractToken, false)
			// Получим refreshToken
			token := jwtStorage.GetRefreshTokenById(strconv.FormatInt(accessTokenMetadata.AccountId, 10))
			if refreshToken, err := security.GetRefreshTokenByStr(token, true); err == nil {
				// Проверим, не истек ли он
				if ok := security.TokenValid(refreshToken); ok {
					if newToken, err := security.CreateToken(accessTokenMetadata.AccountId); err == nil {
						// Упакуем
						aCookie := fasthttp.Cookie{}
						aCookie.SetKey(accessTokenKey)
						aCookie.SetValue(newToken.AccessToken)
						ctx.Response.Header.SetCookie(&aCookie)

						jwtStorage.PutTokens(accessTokenMetadata.AccountId, newToken.AccessToken, newToken.RefreshToken)
						ctx.SetUserValue(security.ContextKey().AccountIdKey, accessTokenMetadataExtracted.AccountId)
						next(ctx)
						return
					}
				}
			}
		}
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
	}
}

func (m authMiddleWare) WareCommon(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		next(ctx)
		m.wareAll(ctx)
	}
}

func (m authMiddleWare) WareLogin(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer m.wareAll(ctx)

		type loginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var lReq loginRequest
		if err := json.Unmarshal(ctx.Request.Body(), &lReq); err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		acc, err := m.service.AccountService.GetByEmail(lReq.Email)
		if err == nil && m.passwordHash.ComparePassword(*acc.Password, lReq.Password) {
			jwtStorage := security.JwtStorage()
			if ok := jwtStorage.IsExistAccessTokenById(acc.Id); !ok {
				token, err := security.CreateToken(acc.Id)
				if err != nil {
					ctx.SetStatusCode(http.StatusUnprocessableEntity)
					return
				}

				aCookie := fasthttp.Cookie{}
				aCookie.SetKey(accessTokenKey)
				aCookie.SetValue(token.AccessToken)

				security.JwtStorage().PutTokens(acc.Id, token.AccessToken, token.RefreshToken)

				ctx.Response.Header.SetCookie(&aCookie)
			}
		} else {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		}
	}
}

func (m authMiddleWare) WareLogout(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		request := &ctx.Request
		jwtStorage := security.JwtStorage()
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		if accessTokenMetadataExtracted, err := security.ExtractAccessTokenMetadataExtracted(request, false); err == nil {
			if ok := jwtStorage.DeleteTokenById(strconv.FormatInt(accessTokenMetadataExtracted.AccountId, 10)); ok {
				ctx.SetStatusCode(fasthttp.StatusOK)
			}
		}
		m.wareAll(ctx)
	}
}

func (m authMiddleWare) wareAll(outerCtx *fasthttp.RequestCtx) {
	outerCtx.SetContentType(contentTypeJson)
}

func (m authMiddleWare) WareRegistry(next fasthttp.RequestHandler, security *jwt.Security) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var account domain.Account
		if err := json.Unmarshal(ctx.Request.Body(), &account); err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		hash, err := m.passwordHash.CreateHash(*account.Password)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		account.Password = &hash

		accountId, err := m.service.AccountService.Create(account)
		if err == nil {
			userId, err := m.service.UserService.Create(domain.User{})
			if err == nil {
				userAccountId, err := m.service.UserAccountService.Create(domain.UserAccount{
					UserId:    &userId,
					AccountId: &accountId,
				})
				if err != nil {
					go func() {
						// Не проверяем, что удалилось
						_, _ = m.service.UserAccountService.DeleteById(userAccountId)
						_, _ = m.service.UserService.DeleteById(userId)
						_, _ = m.service.AccountService.DeleteById(accountId)
					}()
				}
			}
		}
		m.wareAll(ctx)
	}
}
