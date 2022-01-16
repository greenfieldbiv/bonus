package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
)

type AccountMarketServiceImpl struct {
	repository *repository.Repository
}

func NewAccountMarketService(repository *repository.Repository) *AccountMarketServiceImpl {
	return &AccountMarketServiceImpl{repository: repository}
}

func (aa AccountMarketServiceImpl) Create(accountMarket domain.AccountMarket) (int64, error) {
	return aa.repository.AccountMarketRepository.Create(accountMarket)
}

func (aa AccountMarketServiceImpl) DeleteById(id int64) (bool, error) {
	return aa.repository.AccountMarketRepository.DeleteById(id)
}

func (aa AccountMarketServiceImpl) IsMarketInAccount(marketId *int64, accountId *int64) (bool, error) {
	return aa.repository.AccountMarketRepository.IsMarketInAccount(marketId, accountId)
}
