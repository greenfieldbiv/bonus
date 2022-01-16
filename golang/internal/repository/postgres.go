package repository

import (
	"bivbonus/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(cfg config.DataBaseConfig) (*sqlx.DB, error) {

	connect, err := sqlx.Connect(cfg.Drivername,
		fmt.Sprintf(
			"host=%s user=%s  password=%s dbname=%s port=%s sslmode=%s",
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.Dbname,
			cfg.Port,
			cfg.Sslmode,
		),
	)
	if err != nil {
		return nil, err
	}
	if err = connect.Ping(); err != nil {
		return nil, err
	}
	return connect, nil
}
