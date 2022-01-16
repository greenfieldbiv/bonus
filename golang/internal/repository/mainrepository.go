package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	Create(account domain.Account) (int64, error)
	GetById(id int64) (*domain.Account, error)
	All() ([]*domain.Account, error)
	IsExist(account int64) (bool, error)
	DeleteById(id int64) (bool, error)
	GetByEmail(email string) (*domain.Account, error)
}

type UserRepository interface {
	Create(user domain.User) (int64, error)
	GetById(id int64) (*domain.User, error)
	All() ([]*domain.User, error)
	IsExistById(id *int64) (bool, error)
	DeleteById(id int64) (bool, error)
}

type TeamRepository interface {
	Create(team domain.Team) (int64, error)
	GetById(id int64) (*domain.Team, error)
	All() ([]*domain.Team, error)
	IsExistBySysName(sysName *string) (bool, error)
	IsExistById(id *int64) (bool, error)
}

type ServiceRepository interface {
	Create(service domain.Service) (int64, error)
	GetById(id int64) (*domain.Service, error)
	All() ([]*domain.Service, error)
	IsExistBySysName(sysName *string) (bool, error)
	IsExistById(id *int64) (bool, error)
}

type AchievementRepository interface {
	Create(achievement domain.Achievement) (int64, error)
	GetById(id int64) (*domain.Achievement, error)
	All() ([]*domain.Achievement, error)
	IsExistBySysName(sysName *string) (bool, error)
	IsExistById(id *int64) (bool, error)
}

type MarketRepository interface {
	Create(market domain.Market) (int64, error)
	GetById(id int64) (*domain.Market, error)
	All() ([]*domain.Market, error)
	IsExistBySysName(sysName *string) (bool, error)
	IsExistById(id *int64) (bool, error)
}

type UserTeamRepository interface {
	Create(userTeam domain.UserTeam) (int64, error)
	IsUserInTeam(teamId *int64, userId *int64) (bool, error)
	GetByUserId(id int64) (*domain.UserTeam, error)
}

type UserAccountRepository interface {
	Create(userAccount domain.UserAccount) (int64, error)
	DeleteById(id int64) (bool, error)
	GetByAccountId(accountId int64) (*domain.UserAccount, error)
}

type AccountAchievementRepository interface {
	Create(userAccount domain.AccountAchievement) (int64, error)
	DeleteById(id int64) (bool, error)
	IsAchievementInAccount(achievementId *int64, accountId *int64) (bool, error)
}

type AccountMarketRepository interface {
	Create(userAccount domain.AccountMarket) (int64, error)
	DeleteById(id int64) (bool, error)
	IsMarketInAccount(marketId *int64, accountId *int64) (bool, error)
}

type Repository struct {
	AccountRepository
	UserRepository
	TeamRepository
	ServiceRepository
	AchievementRepository
	MarketRepository
	UserTeamRepository
	UserAccountRepository
	AccountAchievementRepository
	AccountMarketRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AccountRepository:            NewAccountRepository(db),
		UserRepository:               NewUserRepository(db),
		TeamRepository:               NewTeamRepository(db),
		ServiceRepository:            NewServiceRepository(db),
		AchievementRepository:        NewAchievementRepository(db),
		MarketRepository:             NewMarketRepository(db),
		UserTeamRepository:           NewUserTeamRepository(db),
		UserAccountRepository:        NewUserAccountRepository(db),
		AccountAchievementRepository: NewAccountAchievementRepository(db),
		AccountMarketRepository:      NewAccountMarketRepository(db),
	}
}
