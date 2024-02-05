package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

func (d *Database) Connect(uri string) {
	var err error
	d.DB, err = sqlx.Connect("sqlite", uri)

	if err != nil {
		log.Fatalln(err)
	}
}
