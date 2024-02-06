package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sqlx.DB
}

func (d *Database) Connect(uri string) error {
	var err error

	// TODO: Update to pgsql
	d.DB, err = sqlx.Connect("sqlite3", uri)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	return nil
}
