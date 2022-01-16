package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserAccountRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserAccountRepository(db *sqlx.DB) *UserAccountRepositoryImpl {
	return &UserAccountRepositoryImpl{db: db}
}

func (ua UserAccountRepositoryImpl) Create(userAccount domain.UserAccount) (int64, error) {
	row := ua.db.QueryRow("insert into user_account (userid, accountid) values ($1,$2) returning id",
		userAccount.UserId,
		userAccount.AccountId,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ua UserAccountRepositoryImpl) DeleteById(id int64) (bool, error) {
	_, err := ua.db.Exec("delete from user_account where id = $1", id)
	return err == nil, err
}

func (ua UserAccountRepositoryImpl) GetByAccountId(accountId int64) (*domain.UserAccount, error) {
	var userAccount domain.UserAccount
	if err := ua.db.Get(&userAccount, "select * from user_account where accountid = $1", accountId); err != nil {
		return nil, err
	}
	return &userAccount, nil
}
