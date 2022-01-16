package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserTeamRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserTeamRepository(db *sqlx.DB) *UserTeamRepositoryImpl {
	return &UserTeamRepositoryImpl{db: db}
}

func (u *UserTeamRepositoryImpl) Create(userTeam domain.UserTeam) (int64, error) {
	row := u.db.QueryRow("insert into user_team (teamid, userid) values ($1,$2) returning id",
		userTeam.TeamId,
		userTeam.UserId,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserTeamRepositoryImpl) IsUserInTeam(teamId *int64, userId *int64) (bool, error) {
	var count int64
	if err := u.db.Get(&count,
		"SELECT count(*) from user_team where teamid = $1 and userid = $2",
		teamId,
		userId,
	); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (u *UserTeamRepositoryImpl) GetByUserId(accountId int64) (*domain.UserTeam, error) {
	var userTeam domain.UserTeam
	rows, err := u.db.Queryx("select * from user_team where userid = $1", accountId)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	if err := rows.StructScan(&userTeam); err != nil {
		return nil, err
	}
	return &userTeam, nil
}
