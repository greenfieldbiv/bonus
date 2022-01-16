package domain

type UserTeam struct {
	Id     *int64 `json:"id" db:"id"`
	UserId *int64 `json:"userid" db:"userid"`
	TeamId *int64 `json:"teamid" db:"teamid"`
}
