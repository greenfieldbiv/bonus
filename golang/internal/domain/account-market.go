package domain

type AccountMarket struct {
	Id        *int64 `json:"id" db:"id"`
	AccountId *int64 `json:"accountid" db:"accountid"`
	MarketId  *int64 `json:"marketid" db:"marketid"`
}
