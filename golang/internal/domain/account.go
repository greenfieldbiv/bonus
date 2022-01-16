package domain

import "time"

// Account учетная запись
type Account struct {
	Id        int64     `db:"id"`
	Login     *string   `json:"login"`
	Password  *string   `json:"password,omitempty"`
	Email     *string   `json:"email,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
	Createdat time.Time `db:"createat"`
	Blocked   bool      `db:"blocked"`
}
