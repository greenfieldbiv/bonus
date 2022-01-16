package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
	"errors"
)

type MarketServiceImpl struct {
	repository *repository.Repository
	service    Service
}

func NewMarketService(repository *repository.Repository) *MarketServiceImpl {
	return &MarketServiceImpl{repository: repository}
}

func (m MarketServiceImpl) GetMarketInfo(id int64) (*domain.Market, error) {
	return m.repository.MarketRepository.GetById(id)
}

func (m MarketServiceImpl) Create(market domain.Market) (int64, error) {
	exist, err := m.IsExistBySysName(market.Sysname)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, errors.New("Market already exist")
	}
	return m.repository.MarketRepository.Create(market)
}

func (m MarketServiceImpl) All() ([]*domain.Market, error) {
	return m.repository.MarketRepository.All()
}

func (m MarketServiceImpl) IsExistBySysName(sysName *string) (bool, error) {
	return m.repository.MarketRepository.IsExistBySysName(sysName)
}

func (m MarketServiceImpl) IsExistById(id *int64) (bool, error) {
	return m.repository.MarketRepository.IsExistById(id)
}

func (m MarketServiceImpl) setService(service Service) {
	m.service = service
}
