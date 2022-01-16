package jwt

import (
	"bivbonus/pkg/config"
	storage "bivbonus/pkg/security/jwt/storage"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

const (
	accessSecretKey  = "ACCESS_SECRET"
	refreshSecretKey = "REFRESH_SECRET"
	accountIdKey     = "account_id"
)

type ContextKey struct {
	AccountIdKey string
}

type Security struct {
	jwtConfig  config.JwtConfig
	jwtStorage storage.JwtStorage
	contextKey ContextKey
}

func (s *Security) ContextKey() ContextKey {
	return s.contextKey
}

func (s *Security) JwtConfig() config.JwtConfig {
	return s.jwtConfig
}

func (s *Security) JwtStorage() storage.JwtStorage {
	return s.jwtStorage
}

func NewSecurity(jwtConfig config.JwtConfig, jwtStorage storage.JwtStorage) *Security {
	err := os.Setenv(accessSecretKey, jwtConfig.Accesssecret)
	if err != nil {
		return nil
	}
	err = os.Setenv(refreshSecretKey, jwtConfig.Refreshsecret)
	if err != nil {
		return nil
	}
	return &Security{
		jwtConfig:  jwtConfig,
		jwtStorage: jwtStorage,
		contextKey: ContextKey{AccountIdKey: accountIdKey},
	}
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    float64
	RtExpires    float64
}

type CustomTokenDetails struct {
	RefreshToken *jwt.Token
	AccessToken  *jwt.Token
	AccountId    int64
}

func (s *Security) CreateToken(accountId int64) (*TokenDetails, error) {
	var err error

	td := &TokenDetails{}
	td.AtExpires = float64(time.Now().Add(time.Minute * 15).Unix())
	v4, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	td.AccessUuid = v4.String()

	td.RtExpires = float64(time.Now().Add(time.Hour * 24 * 7).Unix())
	newV4, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	td.RefreshUuid = newV4.String()
	//Creating Access Token
	err = s.createAccessToken(accountId, td)
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	err = s.createRefreshToken(accountId, td)
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (s *Security) createAccessToken(accountId int64, td *TokenDetails) error {
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.AccessUuid
	atClaims[accountIdKey] = accountId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	signedString, err := at.SignedString([]byte(os.Getenv(accessSecretKey)))
	if err == nil {
		td.AccessToken = signedString
	}
	return err
}

func (s *Security) createRefreshToken(accountId int64, td *TokenDetails) error {
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims[accountIdKey] = accountId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	signedString, err := rt.SignedString([]byte(os.Getenv(refreshSecretKey)))
	if err == nil {
		td.RefreshToken = signedString
	}
	return err
}

func (s *Security) ExtractToken(r *fasthttp.Request) string {
	bearToken := string(r.Header.Peek("Authorization"))
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s *Security) verifyAccessToken(r *fasthttp.Request) (*jwt.Token, error) {
	return s.GetAccessTokenByStr(s.ExtractToken(r), true)
}

func (s *Security) verifyRefreshToken(r *fasthttp.Request) (*jwt.Token, error) {
	return s.GetRefreshTokenByStr(s.ExtractToken(r), true)
}

func (s *Security) GetAccessTokenByStr(tokenString string, isVerifyClaims bool) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(accessSecretKey)), nil
	})
	if !isVerifyClaims {
		err = nil
	}
	return token, err
}

func (s *Security) GetRefreshTokenByStr(jwtToken string, isVerifyClaims bool) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(refreshSecretKey)), nil
	})
	if !isVerifyClaims {
		err = nil
	}
	return token, err
}

func (s *Security) TokenValid(token *jwt.Token) bool {
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), true) {
		return true
	}
	return false
}

func (s *Security) TokenValidByRequest(r *fasthttp.Request) error {
	token, err := s.verifyAccessToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid && !token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), true) {
		return err
	}
	return nil
}

func (s *Security) ExtractAccessTokenMetadata(token string, isDoValidation bool) (*CustomTokenDetails, error) {
	jwtToken, err := s.GetAccessTokenByStr(token, isDoValidation)
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid || !isDoValidation {
		accountId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims[accountIdKey]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &CustomTokenDetails{
			AccessToken: jwtToken,
			AccountId:   accountId,
		}, nil
	}
	return nil, err
}

func (s *Security) ExtractAccessTokenMetadataExtracted(r *fasthttp.Request, isDoValidation bool) (*CustomTokenDetails, error) {
	jwtToken, err := s.GetAccessTokenByStr(s.ExtractToken(r), isDoValidation)
	if jwtToken == nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid || !isDoValidation {
		accountId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims[accountIdKey]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &CustomTokenDetails{
			AccessToken: jwtToken,
			AccountId:   accountId,
		}, nil
	}
	return nil, err
}

func (s *Security) ExtractRefreshTokenMetadata(token string, isDoValidation bool) (*CustomTokenDetails, error) {
	jwtToken, err := s.GetRefreshTokenByStr(token, isDoValidation)
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid || !isDoValidation {
		accountId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims[accountIdKey]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &CustomTokenDetails{
			RefreshToken: jwtToken,
			AccountId:    accountId,
		}, nil
	}
	return nil, err
}

func (s *Security) ExtractRefreshTokenMetadataExtracted(r *fasthttp.Request, isDoValidation bool) (*CustomTokenDetails, error) {
	jwtToken, err := s.GetRefreshTokenByStr(s.ExtractToken(r), isDoValidation)
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid || !isDoValidation {
		accountId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims[accountIdKey]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &CustomTokenDetails{
			RefreshToken: jwtToken,
			AccountId:    accountId,
		}, nil
	}
	return nil, err
}
