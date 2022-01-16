package domain

type Market struct {
	Id      *int64   `json:"id" db:"id"`
	Name    *string  `json:"name" db:"name"`
	Sysname *string  `json:"sysname" db:"sysname"`
	Level   *int     `json:"level" db:"level"`
	Cost    *float32 `json:"cost" db:"cost"`
}
