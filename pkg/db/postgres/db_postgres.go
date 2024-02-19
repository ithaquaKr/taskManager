package postgres

import (
	"time"

	"github.com/ithaquaKr/taskManager/config"
	_ "github.com/jackc/pgx/v4/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConn     = 10
	connMaxLifeTime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPostgresConn(c *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(c.Postgres.PostgresPgDriver, c.Postgres.PostgresURL)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConn)
	db.SetConnMaxLifetime(connMaxLifeTime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
