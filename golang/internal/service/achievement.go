package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
	"errors"
)

type AchievementServiceImpl struct {
	repository *repository.Repository
	service    Service
}

func NewAchievementService(repository *repository.Repository) *AchievementServiceImpl {
	return &AchievementServiceImpl{
		repository: repository,
	}
}

func (a AchievementServiceImpl) GetAchievementInfo(id int64) (*domain.Achievement, error) {
	return a.repository.AchievementRepository.GetById(id)
}

func (a AchievementServiceImpl) Create(achievement domain.Achievement) (int64, error) {
	exist, err := a.IsExistBySysName(achievement.Sysname)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, errors.New("Achievement already exist")
	}
	return a.repository.AchievementRepository.Create(achievement)
}

func (a AchievementServiceImpl) All() ([]*domain.Achievement, error) {
	return a.repository.AchievementRepository.All()
}

func (a AchievementServiceImpl) IsExistBySysName(sysName *string) (bool, error) {
	return a.repository.AchievementRepository.IsExistBySysName(sysName)
}

func (a AchievementServiceImpl) setService(service Service) {
	a.service = service
}
