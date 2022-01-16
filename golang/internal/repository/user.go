package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryIml struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepositoryIml {
	return &UserRepositoryIml{
		db: db,
	}
}

func (u *UserRepositoryIml) Create(user domain.User) (int64, error) {
	row := u.db.QueryRow(
		"INSERT INTO \"user\" (name, surname, patronymic) values ($1, $2, $3) returning id",
		user.Name,
		user.Surname,
		user.Patronymic,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepositoryIml) GetById(id int64) (*domain.User, error) {
	var user domain.User
	rows, err := u.db.Queryx("SELECT * from \"user\" where id = $1", id)
	if err := err; err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	if err = rows.StructScan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryIml) All() ([]*domain.User, error) {
	rows, err := u.db.Query("SELECT * from \"user\"")
	defer rows.Close()
	if err != nil {
		return []*domain.User{}, err
	}
	var uList []*domain.User
	for rows.Next() {
		var (
			id         int64
			name       *string
			surname    *string
			patronymic *string
		)
		err := rows.Scan(&id, &name, &surname, &patronymic)
		if err != nil {
			return nil, err
		}
		t := &domain.User{
			Name:       name,
			Surname:    surname,
			Patronymic: patronymic,
		}
		uList = append(uList, t)
	}

	return uList, nil
}

func (u *UserRepositoryIml) IsExistById(id *int64) (bool, error) {
	var count int64
	if err := u.db.Get(&count, "SELECT count(*) from \"user\" where id = $1", id); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (u *UserRepositoryIml) DeleteById(id int64) (bool, error) {
	_, err := u.db.Exec("delete from \"user\" where id = $1", id)
	return err == nil, err
}
