package domain

type UserAccount struct {
	Id        *int64 `json:"id" db:"id"`
	UserId    *int64 `json:"userid" db:"userid"`
	AccountId *int64 `json:"accountid" db:"accountid"`
}
