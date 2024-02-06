package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	DB  *sqlx.DB
	url string
}

func (d *Database) Connect() error {
	var err error

	// TODO: Update to pgsql
	d.DB, err = sqlx.Connect("postgres", d.url)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	return nil
}
