package handler

import (
	"bivbonus/internal/common"
	"bivbonus/internal/service"
	"bivbonus/pkg/security/jwt"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type accountInfoResponse struct {
	Login     *string   `json:"login"`
	Email     *string   `json:"email,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
	CreatedAt time.Time `json:"createdat"`
	Blocked   bool      `json:"blocked"`
}

type AccountHandler struct {
	BaseHandler
}

func NewAccountHandler(service *service.Service, security *jwt.Security) *AccountHandler {
	return &AccountHandler{
		BaseHandler{
			service: service, security: security,
		},
	}
}

func (a *AccountHandler) GetAccountInfo(ctx *fasthttp.RequestCtx) {
	pathParamId, err := strconv.ParseInt((ctx.Value("id")).(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	acInfo, err := a.service.AccountService.GetAccountInfo(pathParamId)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	accountInfoResponse := accountInfoResponse{
		Login:     acInfo.Login,
		Email:     acInfo.Email,
		Phone:     acInfo.Phone,
		CreatedAt: acInfo.Createdat,
		Blocked:   acInfo.Blocked,
	}
	commonResponse := common.CommonResponse{
		Data: accountInfoResponse,
	}
	response, err := json.Marshal(commonResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(response)
}

func (a *AccountHandler) All(ctx *fasthttp.RequestCtx) {
	all, err := a.service.AccountService.All()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&all)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

type addAchieveToAccountRequest struct {
	AccountId     int64 `json:"accountid"`
	AchievementId int64 `json:"achievementid"`
}

type addAchieveToAccountResponse struct {
	Id int64 `json:"id"`
}

func (a AccountHandler) AddAchievement(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	var addAchieveToAccountRequest addAchieveToAccountRequest
	if err := json.Unmarshal(body, &addAchieveToAccountRequest); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	id, err := a.service.AccountService.AddAchievement(
		addAchieveToAccountRequest.AccountId,
		addAchieveToAccountRequest.AchievementId,
	)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	aUr := addAchieveToAccountResponse{Id: id}
	response, err := json.Marshal(aUr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

type addMarketToAccountRequest struct {
	AccountId int64 `json:"accountid"`
	MarketId  int64 `json:"marketid"`
}

type addMarketToAccountResponse struct {
	Id int64 `json:"id"`
}

func (a *AccountHandler) AddMarket(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	var addMarketToAccountRequest addMarketToAccountRequest
	if err := json.Unmarshal(body, &addMarketToAccountRequest); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	id, err := a.service.AccountService.AddMarket(
		addMarketToAccountRequest.AccountId,
		addMarketToAccountRequest.MarketId,
	)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	aUr := addMarketToAccountResponse{Id: id}
	response, err := json.Marshal(aUr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

type getAccountFullInfoUser struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
}

type getAccountFullInfoTeam struct {
	Name    *string `json:"name"`
	Sysname *string `json:"sysname"`
}

type getAccountFullInfoService struct {
	Name    *string `json:"name"`
	Sysname *string `json:"sysname"`
}

type getAccountFullInfoResponse struct {
	Email    *string                      `json:"email,omitempty"`
	Phone    *string                      `json:"phone,omitempty"`
	User     *getAccountFullInfoUser      `json:"user"`
	Team     *getAccountFullInfoTeam      `json:"team"`
	Services []*getAccountFullInfoService `json:"services"`
}

func (a *AccountHandler) GetAccountFullInfo(ctx *fasthttp.RequestCtx) {
	accountId := ctx.UserValue(a.security.ContextKey().AccountIdKey).(int64)
	// Загружаем аккаунт
	account, err := a.service.AccountService.GetAccountInfo(accountId)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	accountFullInfoResponse := getAccountFullInfoResponse{
		Email: account.Email,
		Phone: account.Phone,
	}

	// Загружаем пользователя
	userAccount, err := a.service.UserAccountService.GetUserAccountByAccountId(account.Id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	if userAccount != nil && userAccount.UserId != nil {
		user, err := a.service.UserService.GetUserInfo(*userAccount.UserId)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		if user != nil && user.Id != nil {
			accountFullInfoResponse.User = &getAccountFullInfoUser{
				Name:       user.Name,
				Surname:    user.Surname,
				Patronymic: user.Patronymic,
			}
			// Загружаем команду
			userTeam, err := a.service.UserTeamService.GetUserTeamByUserId(*user.Id)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
				return
			}
			if userTeam != nil && userTeam.UserId != nil {
				team, err := a.service.TeamService.GetTeamInfo(*userTeam.TeamId)
				if err != nil {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					return
				}
				if team != nil {
					accountFullInfoResponse.Team = &getAccountFullInfoTeam{
						Name:    team.Name,
						Sysname: team.Sysname,
					}
				}
			} else {
				accountFullInfoResponse.Team = nil
			}
			// Загружаем достижения
			accountFullInfoResponse.Services = nil
		} else {
			accountFullInfoResponse.User = nil
		}
	}

	response, err := json.Marshal(accountFullInfoResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}
