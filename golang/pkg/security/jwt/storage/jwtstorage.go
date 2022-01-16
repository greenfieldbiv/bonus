package storage

type JwtStorage interface {
	GetAccessToken(id int64) string
	GetRefreshToken(id int64) string
	GetRefreshTokenByAccessToken(accessToken string) string
	GetRefreshTokenById(id string) string
	PutTokens(id int64, access string, refresh string) bool
	PutAccessToken(id int64, access string) bool
	PutRefreshToken(id int64, refresh string) bool
	DeleteAccessToken(id int64, access string) bool
	DeleteRefreshToken(id int64, refresh string) bool
	DeleteTokenById(id string) bool
	IsExistAccessToken(id int64, access string) bool
	IsExistRefreshToken(id int64, refresh string) bool
	IsExistAccessTokenById(id int64) bool
	IsExistRefreshTokenById(id int64) bool
}
