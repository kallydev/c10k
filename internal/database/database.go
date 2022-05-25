package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/kallydev/c10k/internal/database/model"
	_ "github.com/lib/pq"
)

var (
	createTableFuncs = []func(db *sqlx.DB) error{
		model.CreatePostTable,
	}
)

func Dial(dsn string) (db *sqlx.DB, err error) {
	if db, err = sqlx.Open("postgres", dsn); err != nil {
		return nil, err
	}

	if _, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		return nil, err
	}

	for _, createTableFunc := range createTableFuncs {
		if err := createTableFunc(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}
