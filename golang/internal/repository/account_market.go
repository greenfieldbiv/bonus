package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AccountMarketRepositoryImpl struct {
	db *sqlx.DB
}

func NewAccountMarketRepository(db *sqlx.DB) *AccountMarketRepositoryImpl {
	return &AccountMarketRepositoryImpl{db: db}
}

func (a AccountMarketRepositoryImpl) Create(accountMarket domain.AccountMarket) (int64, error) {
	row := a.db.QueryRow("insert into account_market (marketid, accountid) values ($1,$2) returning id",
		accountMarket.MarketId,
		accountMarket.AccountId,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a AccountMarketRepositoryImpl) DeleteById(id int64) (bool, error) {
	_, err := a.db.Exec("delete from account_market where id = $1", id)
	return err == nil, err
}

func (a AccountMarketRepositoryImpl) IsMarketInAccount(marketId *int64, accountId *int64) (bool, error) {
	var count int64
	if err := a.db.Get(&count,
		"SELECT count(*) from account_market where marketid = $1 and accountid = $2",
		marketId,
		accountId,
	); err != nil {
		return false, err
	}
	return count != 0, nil
}
