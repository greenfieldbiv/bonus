package service

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/repository"
)

type AccountService interface {
	Create(account domain.Account) (id int64, err error)
	GetAccountInfo(id int64) (*domain.Account, error)
	All() ([]*domain.Account, error)
	IsExist(id int64) (bool, error)
	GetByEmail(email string) (*domain.Account, error)
	DeleteById(id int64) (bool, error)
	AddAchievement(accountId int64, achievementId int64) (id int64, err error)
	AddMarket(accountId int64, marketId int64) (id int64, err error)
}

type UserService interface {
	GetUserInfo(id int64) (*domain.User, error)
	Create(user domain.User) (int64, error)
	All() ([]*domain.User, error)
	IsExistById(id *int64) (bool, error)
	DeleteById(id int64) (bool, error)
}

type TeamService interface {
	GetTeamInfo(id int64) (*domain.Team, error)
	Create(team domain.Team) (int64, error)
	All() ([]*domain.Team, error)
	IsExistBySysName(sysName *string) (bool, error)
	IsExistById(id *int64) (bool, error)
	AddUser(teamId *int64, userId *int64) (int64, error)
}

type ServiceService interface {
	GetServiceInfo(id int64) (*domain.Service, error)
	Create(team domain.Service) (int64, error)
	All() ([]*domain.Service, error)
	IsExistBySysName(sysName *string) (bool, error)
}

type AchievementService interface {
	GetAchievementInfo(id int64) (*domain.Achievement, error)
	Create(team domain.Achievement) (int64, error)
	All() ([]*domain.Achievement, error)
	IsExistBySysName(sysName *string) (bool, error)
}

type MarketService interface {
	GetMarketInfo(id int64) (*domain.Market, error)
	Create(team domain.Market) (int64, error)
	All() ([]*domain.Market, error)
	IsExistBySysName(sysName *string) (bool, error)
	IsExistById(id *int64) (bool, error)
}

type UserTeamService interface {
	Create(userTeam domain.UserTeam) (int64, error)
	isUserInTeam(teamId *int64, userId *int64) (bool, error)
	GetUserTeamByUserId(id int64) (*domain.UserTeam, error)
}

type UserAccountService interface {
	Create(userAccount domain.UserAccount) (int64, error)
	DeleteById(id int64) (bool, error)
	GetUserAccountByAccountId(accountId int64) (*domain.UserAccount, error)
}

type AccountAchievementService interface {
	Create(userAccount domain.AccountAchievement) (int64, error)
	DeleteById(id int64) (bool, error)
	IsAchievementInAccount(achievementId *int64, accountId *int64) (bool, error)
}

type AccountMarketService interface {
	Create(userAccount domain.AccountMarket) (int64, error)
	DeleteById(id int64) (bool, error)
	IsMarketInAccount(marketId *int64, accountId *int64) (bool, error)
}

type Service struct {
	AccountService
	UserService
	TeamService
	ServiceService
	AchievementService
	MarketService
	UserTeamService
	UserAccountService
	AccountAchievementService
	AccountMarketService
}

func NewService(repository *repository.Repository) *Service {
	accountService := NewAccountService(repository)
	userService := NewUserService(repository)
	teamService := NewTeamService(repository)
	serviceService := NewServiceService(repository)
	achievementService := NewAchievementService(repository)
	marketService := NewMarketService(repository)
	userTeamService := NewUserTeamService(repository)
	userAccountService := NewUserAccountService(repository)
	accountAchievementService := NewAccountAchievementService(repository)
	accountMarketService := NewAccountMarketService(repository)

	service := Service{
		AccountService:            accountService,
		UserService:               userService,
		TeamService:               teamService,
		ServiceService:            serviceService,
		AchievementService:        achievementService,
		MarketService:             marketService,
		UserTeamService:           userTeamService,
		UserAccountService:        userAccountService,
		AccountAchievementService: accountAchievementService,
		AccountMarketService:      accountMarketService,
	}

	accountService.setService(service)
	teamService.setService(service)

	return &service
}
