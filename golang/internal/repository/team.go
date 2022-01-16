package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type TeamRepositoryIml struct {
	db *sqlx.DB
}

func NewTeamRepository(db *sqlx.DB) *TeamRepositoryIml {
	return &TeamRepositoryIml{
		db: db,
	}
}

func (u *TeamRepositoryIml) Create(team domain.Team) (int64, error) {
	row := u.db.QueryRow(
		"INSERT INTO team (name, sysname) values ($1, $2) returning id",
		team.Name,
		team.Sysname,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *TeamRepositoryIml) GetById(id int64) (*domain.Team, error) {
	var team domain.Team
	if err := u.db.Get(&team, "SELECT * from team where id = $1", id); err != nil {
		return nil, err
	}
	return &team, nil
}

func (u *TeamRepositoryIml) All() ([]*domain.Team, error) {
	rows, err := u.db.Query("SELECT * from team")
	defer rows.Close()
	if err != nil {
		return []*domain.Team{}, err
	}
	var tList []*domain.Team
	for rows.Next() {
		var (
			id      int64
			name    *string
			sysname *string
		)
		err := rows.Scan(&id, &name, &sysname)
		if err != nil {
			return nil, err
		}
		t := &domain.Team{
			Name:    name,
			Sysname: sysname,
		}
		tList = append(tList, t)
	}

	return tList, nil
}

func (u *TeamRepositoryIml) IsExistBySysName(sysName *string) (bool, error) {
	var count int64
	if err := u.db.Get(&count, "SELECT count(*) from team where sysname = $1", sysName); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (u *TeamRepositoryIml) IsExistById(id *int64) (bool, error) {
	var count int64
	if err := u.db.Get(&count, "SELECT count(*) from team where id = $1", id); err != nil {
		return false, err
	}
	return count != 0, nil
}
