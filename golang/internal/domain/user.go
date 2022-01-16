package domain

type User struct {
	Id         *int64  `db:"id"`
	Name       *string `db:"name"`
	Surname    *string `db:"surname"`
	Patronymic *string `db:"patronymic"`
}
