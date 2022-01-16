package domain

type Team struct {
	Id      *int64  `json:"id" db:"id"`
	Name    *string `json:"name" db:"name"`
	Sysname *string `json:"sysname" db:"sysname"`
}
