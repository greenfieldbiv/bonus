package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
)

type UserAccountServiceImpl struct {
	repository *repository.Repository
}

func NewUserAccountService(repository *repository.Repository) *UserAccountServiceImpl {
	return &UserAccountServiceImpl{repository: repository}
}

func (ua UserAccountServiceImpl) Create(userAccount domain.UserAccount) (int64, error) {
	return ua.repository.UserAccountRepository.Create(userAccount)
}

func (ua UserAccountServiceImpl) DeleteById(id int64) (bool, error) {
	_, err := ua.repository.UserAccountRepository.DeleteById(id)
	return err == nil, err
}

func (ua UserAccountServiceImpl) GetUserAccountByAccountId(accountId int64) (*domain.UserAccount, error) {
	return ua.repository.UserAccountRepository.GetByAccountId(accountId)
}
