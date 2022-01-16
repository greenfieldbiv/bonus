package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
	"errors"
)

type ServiceServiceImpl struct {
	repository *repository.Repository
	service    Service
}

func NewServiceService(repository *repository.Repository) *ServiceServiceImpl {
	return &ServiceServiceImpl{repository: repository}
}

func (s *ServiceServiceImpl) GetServiceInfo(id int64) (*domain.Service, error) {
	return s.repository.ServiceRepository.GetById(id)
}

func (s *ServiceServiceImpl) Create(service domain.Service) (int64, error) {
	exist, err := s.IsExistBySysName(service.SysName)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, errors.New("Service already exist")
	}
	return s.repository.ServiceRepository.Create(service)
}

func (s *ServiceServiceImpl) All() ([]*domain.Service, error) {
	return s.repository.ServiceRepository.All()
}

func (s *ServiceServiceImpl) IsExistBySysName(sysName *string) (bool, error) {
	return s.repository.ServiceRepository.IsExistBySysName(sysName)
}

func (s *ServiceServiceImpl) setService(service Service) {
	s.service = service
}
