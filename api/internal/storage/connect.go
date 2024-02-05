package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sqlx.DB
}

func (d *Database) Connect(uri string) error {
	var err error
	d.DB, err = sqlx.Connect("sqlite3", uri)

	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	return nil
}
