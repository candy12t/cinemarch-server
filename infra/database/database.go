package database

import (
	"database/sql"

	"github.com/candy12t/cinema-search-server/config"
)

func NewDB() (*sql.DB, func() error, error) {
	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, db.Close, err
	}

	return db, db.Close, nil
}
