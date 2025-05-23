package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config interface {
	GetDSN() string // Data Source Name
	GetMaxOpenConns() int
	GetMaxIdleConns() int
	GetConnMaxIdleTime() time.Duration
	GetConnMaxLifetime() time.Duration
}

var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func ConnectDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.GetDSN())

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.GetMaxOpenConns())
	db.SetMaxIdleConns(cfg.GetMaxIdleConns())
	db.SetConnMaxIdleTime(cfg.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.GetConnMaxLifetime())

	return db, nil
}
