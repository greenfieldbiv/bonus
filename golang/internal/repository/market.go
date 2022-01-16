package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type MarketRepositoryImpl struct {
	db *sqlx.DB
}

func NewMarketRepository(db *sqlx.DB) *MarketRepositoryImpl {
	return &MarketRepositoryImpl{db: db}
}

func (m MarketRepositoryImpl) Create(market domain.Market) (int64, error) {
	row := m.db.QueryRow("insert into market (name ,sysname ,level ,cost) values ($1,$2,$3,$4) returning id",
		market.Name,
		market.Sysname,
		market.Level,
		market.Cost,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (m MarketRepositoryImpl) GetById(id int64) (*domain.Market, error) {
	var market domain.Market
	if err := m.db.Get(&market, "SELECT * from market where id = $1", id); err != nil {
		return nil, err
	}
	return &market, nil
}

func (m MarketRepositoryImpl) All() ([]*domain.Market, error) {
	rows, err := m.db.Query("SELECT * from market")
	defer rows.Close()
	if err != nil {
		return []*domain.Market{}, err
	}
	var mList []*domain.Market
	for rows.Next() {
		var (
			id      int64
			name    *string
			sysname *string
			level   *int
			cost    *float32
		)
		err := rows.Scan(&id, &name, &sysname, &level, &cost)
		if err != nil {
			return nil, err
		}
		t := &domain.Market{
			Name:    name,
			Sysname: sysname,
			Level:   level,
			Cost:    cost,
		}
		mList = append(mList, t)
	}

	return mList, nil
}

func (m MarketRepositoryImpl) IsExistBySysName(sysName *string) (bool, error) {
	var count int64
	if err := m.db.Get(&count, "SELECT count(*) from market where sysname = $1", sysName); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (m MarketRepositoryImpl) IsExistById(id *int64) (bool, error) {
	var count int64
	if err := m.db.Get(&count, "SELECT count(*) from market where id = $1", id); err != nil {
		return false, err
	}
	return count != 0, nil
}
