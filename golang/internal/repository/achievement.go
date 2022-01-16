package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AchievementRepositoryImpl struct {
	db *sqlx.DB
}

func NewAchievementRepository(db *sqlx.DB) *AchievementRepositoryImpl {
	return &AchievementRepositoryImpl{db: db}
}

func (a AchievementRepositoryImpl) Create(achievement domain.Achievement) (int64, error) {
	row := a.db.QueryRow("insert into achievement (name ,sysname ,level ,cost) values ($1,$2,$3,$4) returning id",
		achievement.Name,
		achievement.Sysname,
		achievement.Level,
		achievement.Cost,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a AchievementRepositoryImpl) GetById(id int64) (*domain.Achievement, error) {
	var ac domain.Achievement
	if err := a.db.Get(&ac, "SELECT * from achievement where id = $1", id); err != nil {
		return nil, err
	}
	return &ac, nil
}

func (a AchievementRepositoryImpl) All() ([]*domain.Achievement, error) {
	rows, err := a.db.Query("SELECT * from achievement")
	defer rows.Close()
	if err != nil {
		return []*domain.Achievement{}, err
	}
	var aList []*domain.Achievement
	for rows.Next() {
		var (
			id      int64
			name    *string
			sysname *string
			level   *int
			cost    *float32
		)
		err := rows.Scan(&id, &name, &sysname, &level, &cost)
		if err != nil {
			return nil, err
		}
		t := &domain.Achievement{
			Name:    name,
			Sysname: sysname,
			Level:   level,
			Cost:    cost,
		}
		aList = append(aList, t)
	}

	return aList, nil
}

func (a AchievementRepositoryImpl) IsExistBySysName(sysName *string) (bool, error) {
	var count int64
	if err := a.db.Get(&count, "SELECT count(*) from achievement where sysName = $1", sysName); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (a AchievementRepositoryImpl) IsExistById(id *int64) (bool, error) {
	var count int64
	if err := a.db.Get(&count, "SELECT count(*) from achievement where id = $1", id); err != nil {
		return false, err
	}
	return count != 0, nil
}
