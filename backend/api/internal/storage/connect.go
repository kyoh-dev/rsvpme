package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	DB           *sqlx.DB
	MaxOpenConns int
	MaxIdleConns int
}

func (d *Database) Init(url string) error {
	var err error
	d.DB, err = sqlx.Connect("postgres", url)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	return nil
}
