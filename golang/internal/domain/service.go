package domain

type Service struct {
	Id      int64    `db:"id"`
	Name    *string  `db:"name"`
	SysName *string  `db:"sysname"`
	Cost    *float32 `db:"cost"`
}
