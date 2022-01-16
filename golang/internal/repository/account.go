package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type AccountRepositoryImpl struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		db: db,
	}
}

func (r AccountRepositoryImpl) Create(account domain.Account) (int64, error) {
	row := r.db.QueryRow(
		"INSERT INTO account (password, email, createat) VALUES ($1, $2, $3) returning id",
		account.Password,
		account.Email,
		account.Createdat,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r AccountRepositoryImpl) GetById(id int64) (*domain.Account, error) {
	var account domain.Account
	if err := r.db.Get(&account, "SELECT * FROM account where id = $1", id); err != nil {
		return nil, err
	}
	return &account, nil
}

func (r AccountRepositoryImpl) All() ([]*domain.Account, error) {
	rows, err := r.db.Query("SELECT * from account")
	defer rows.Close()
	if err != nil {
		return []*domain.Account{}, err
	}
	var acList []*domain.Account
	for rows.Next() {
		var (
			id        int64
			login     *string
			password  *string
			email     *string
			phone     *string
			createdat time.Time
			blocked   bool
		)
		err := rows.Scan(&id, &login, &password, &email, &phone, &createdat, &blocked)
		if err != nil {
			return nil, err
		}
		ac := &domain.Account{
			Login:     login,
			Password:  password,
			Email:     email,
			Phone:     phone,
			Createdat: createdat,
			Blocked:   blocked,
		}
		acList = append(acList, ac)
	}

	return acList, nil
}

func (r AccountRepositoryImpl) DeleteById(id int64) (bool, error) {
	_, err := r.db.Exec("delete from account where id = $1", id)
	return err == nil, err
}

func (r AccountRepositoryImpl) IsExist(id int64) (bool, error) {
	// TODO реализовать после переноса метода из сервиса авторизации в сервис аккаунта!
	return false, nil
}

func (r AccountRepositoryImpl) GetByEmail(email string) (*domain.Account, error) {
	var account domain.Account
	if err := r.db.Get(&account, "SELECT * FROM account where email = $1", email); err != nil {
		return nil, err
	}
	return &account, nil
}
