package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
	"errors"
)

type AccountServiceImpl struct {
	repository *repository.Repository
	service    Service
}

func NewAccountService(repository *repository.Repository) *AccountServiceImpl {
	return &AccountServiceImpl{repository: repository}
}

func (s *AccountServiceImpl) Create(account domain.Account) (int64, error) {
	accountId, err := s.repository.AccountRepository.Create(account)
	if err != nil {
		return 0, err
	}
	userId, err := s.service.UserService.Create(domain.User{})
	if err != nil {
		return 0, err
	}
	userAccountId, err := s.service.UserAccountService.Create(domain.UserAccount{
		UserId:    &userId,
		AccountId: &accountId,
	})
	if err != nil {
		// Удаляем в отдельном потоке
		go func() {
			s.service.UserAccountService.DeleteById(userAccountId)
			s.service.UserService.DeleteById(userId)
			s.service.AccountService.DeleteById(accountId)
		}()
	}
	return userAccountId, nil
}

func (s *AccountServiceImpl) GetAccountInfo(id int64) (*domain.Account, error) {
	return s.repository.AccountRepository.GetById(id)
}

func (s *AccountServiceImpl) All() ([]*domain.Account, error) {
	return s.repository.AccountRepository.All()
}

func (s *AccountServiceImpl) IsExist(id int64) (bool, error) {
	return s.repository.AccountRepository.IsExist(id)
}

func (s *AccountServiceImpl) DeleteById(id int64) (bool, error) {
	return s.repository.AccountRepository.DeleteById(id)
}

func (s *AccountServiceImpl) AddAchievement(accountId int64, achievementId int64) (id int64, err error) {
	isAccountExist, _ := s.service.AccountService.IsExist(accountId)
	if !isAccountExist {
		return 0, errors.New("Не удалось добавить достижение")
	}

	isAchievementInAccount, _ := s.service.AccountAchievementService.IsAchievementInAccount(&achievementId, &accountId)
	if isAchievementInAccount {
		return 0, errors.New("У данного аккаунта уже есть такое достижение")
	}

	return s.repository.AccountAchievementRepository.Create(domain.AccountAchievement{
		AccountId:     &accountId,
		AchievementId: &achievementId,
	})
}

func (s *AccountServiceImpl) AddMarket(accountId int64, marketId int64) (id int64, err error) {
	isAccountExist, _ := s.service.AccountService.IsExist(accountId)
	if !isAccountExist {
		return 0, errors.New("Не удалось добавить позицию магазина")
	}

	isMarketExist, _ := s.service.AccountMarketService.IsMarketInAccount(&marketId, &accountId)
	if isMarketExist {
		return 0, errors.New("У данного аккаунта уже есть такая позиция магазина")
	}

	isMarketInAccount, _ := s.service.AccountMarketService.IsMarketInAccount(&marketId, &accountId)
	if isMarketInAccount {
		return 0, errors.New("Данная позиция уже существует")
	}
	return s.repository.AccountMarketRepository.Create(domain.AccountMarket{
		AccountId: &accountId,
		MarketId:  &marketId,
	})
}

func (s *AccountServiceImpl) setService(service Service) {
	s.service = service
}

func (s *AccountServiceImpl) GetByEmail(email string) (*domain.Account, error) {
	return s.repository.AccountRepository.GetByEmail(email)
}
