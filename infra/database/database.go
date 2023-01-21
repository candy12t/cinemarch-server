package database

import (
	"github.com/candy12t/cinema-search-server/config"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, func() error, error) {
	db, err := sqlx.Open("mysql", config.DSN())
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, db.Close, err
	}

	return db, db.Close, nil
}
