package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
	"errors"
)

type TeamServiceImpl struct {
	repository *repository.Repository
	service    Service
}

func NewTeamService(repository *repository.Repository) *TeamServiceImpl {
	return &TeamServiceImpl{repository: repository}
}

func (t *TeamServiceImpl) GetTeamInfo(id int64) (*domain.Team, error) {
	return t.repository.TeamRepository.GetById(id)
}

func (t *TeamServiceImpl) Create(team domain.Team) (int64, error) {
	exist, err := t.IsExistBySysName(team.Sysname)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, errors.New("Team already exist")
	}
	return t.repository.TeamRepository.Create(team)
}

func (t *TeamServiceImpl) AddUser(teamId *int64, userId *int64) (int64, error) {
	// Проверка на наличие того и другого
	isTeamExist, _ := t.IsExistById(teamId)
	if !isTeamExist {
		return 0, errors.New("Не удалось добавить пользователя в команду")
	}
	//isUserExist, _ := t.service.UserService.IsExistById(userId)
	//if !isUserExist {
	//	return 0, errors.New("Не удалось добавить пользователя в команду")
	//}
	isUserInTeam, _ := t.service.UserTeamService.isUserInTeam(teamId, userId)
	if isUserInTeam {
		return 0, errors.New("Пользователь уже состоит в этой команде")
	}
	return t.repository.UserTeamRepository.Create(domain.UserTeam{
		UserId: userId,
		TeamId: teamId,
	})
}

func (t *TeamServiceImpl) All() ([]*domain.Team, error) {
	return t.repository.TeamRepository.All()
}

func (t *TeamServiceImpl) IsExistBySysName(sysName *string) (bool, error) {
	return t.repository.TeamRepository.IsExistBySysName(sysName)
}

func (t *TeamServiceImpl) IsExistById(id *int64) (bool, error) {
	return t.repository.TeamRepository.IsExistById(id)
}

func (t *TeamServiceImpl) setService(service Service) {
	t.service = service
}
