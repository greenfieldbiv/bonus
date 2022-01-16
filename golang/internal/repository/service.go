package repository

import (
	"bivbonus/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ServiceRepositoryIml struct {
	db *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) *ServiceRepositoryIml {
	return &ServiceRepositoryIml{
		db: db,
	}
}

func (s *ServiceRepositoryIml) Create(service domain.Service) (int64, error) {
	row := s.db.QueryRow(
		"INSERT INTO service (name, sysname, cost) values ($1, $2, $3) returning id",
		service.Name,
		service.SysName,
		service.Cost,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ServiceRepositoryIml) GetById(id int64) (*domain.Service, error) {
	var service domain.Service
	if err := s.db.Get(&service, "SELECT * from service where id = $1", id); err != nil {
		return nil, err
	}
	return &service, nil
}

func (s ServiceRepositoryIml) All() ([]*domain.Service, error) {
	rows, err := s.db.Query("SELECT * from service")
	defer rows.Close()
	if err != nil {
		return []*domain.Service{}, err
	}
	var sList []*domain.Service
	for rows.Next() {
		var (
			id      int64
			name    *string
			sysname *string
			cost    *float32
		)
		err := rows.Scan(&id, &name, &sysname, &cost)
		if err != nil {
			return nil, err
		}
		s := &domain.Service{
			Name:    name,
			SysName: sysname,
			Cost:    cost,
		}
		sList = append(sList, s)
	}

	return sList, nil
}

func (s *ServiceRepositoryIml) IsExistBySysName(sysName *string) (bool, error) {
	var count int64
	if err := s.db.Get(&count, "SELECT count(*) from service where sysname = $1", sysName); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (s *ServiceRepositoryIml) IsExistById(id *int64) (bool, error) {
	var count int64
	if err := s.db.Get(&count, "SELECT count(*) from service where sysname = $1", id); err != nil {
		return false, err
	}
	return count != 0, nil
}
