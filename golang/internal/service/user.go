package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
)

type UserServiceImpl struct {
	repository *repository.Repository
	service    Service
}

func NewUserService(repository *repository.Repository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (u *UserServiceImpl) GetUserInfo(id int64) (*domain.User, error) {
	return u.repository.UserRepository.GetById(id)
}

func (u UserServiceImpl) Create(user domain.User) (int64, error) {
	//exist, err := u.IsExistById(user.Id)
	//if err != nil {
	//	return 0, err
	//}
	//if exist {
	//	return 0, errors.New("User already exist")
	//}
	return u.repository.UserRepository.Create(user)
}

func (u UserServiceImpl) All() ([]*domain.User, error) {
	return u.repository.UserRepository.All()
}

func (u *UserServiceImpl) IsExistById(id *int64) (bool, error) {
	return false, nil
}

func (u *UserServiceImpl) DeleteById(id int64) (bool, error) {
	return u.repository.UserRepository.DeleteById(id)
}

func (u *UserServiceImpl) setService(service Service) {
	u.service = service
}
