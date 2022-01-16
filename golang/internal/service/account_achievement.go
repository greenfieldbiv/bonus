package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
)

type AccountAchievementServiceImpl struct {
	repository *repository.Repository
}

func NewAccountAchievementService(repository *repository.Repository) *AccountAchievementServiceImpl {
	return &AccountAchievementServiceImpl{repository: repository}
}

func (aa AccountAchievementServiceImpl) Create(accountAchievement domain.AccountAchievement) (int64, error) {
	return aa.repository.AccountAchievementRepository.Create(accountAchievement)
}

func (aa AccountAchievementServiceImpl) DeleteById(id int64) (bool, error) {
	return aa.repository.AccountAchievementRepository.DeleteById(id)
}

func (aa AccountAchievementServiceImpl) IsAchievementInAccount(achievementId *int64, accountId *int64) (bool, error) {
	return aa.repository.AccountAchievementRepository.IsAchievementInAccount(achievementId, accountId)
}
