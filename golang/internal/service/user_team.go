package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
)

type UserTeamServiceImpl struct {
	repository *repository.Repository
}

func (u UserTeamServiceImpl) Create(userTeam domain.UserTeam) (int64, error) {
	return u.repository.UserTeamRepository.Create(userTeam)
}

func (u UserTeamServiceImpl) isUserInTeam(teamId *int64, userId *int64) (bool, error) {
	return u.repository.UserTeamRepository.IsUserInTeam(teamId, userId)
}

func (u UserTeamServiceImpl) GetUserTeamByUserId(userId int64) (*domain.UserTeam, error) {
	return u.repository.UserTeamRepository.GetByUserId(userId)
}

func NewUserTeamService(repository *repository.Repository) *UserTeamServiceImpl {
	return &UserTeamServiceImpl{repository: repository}
}
