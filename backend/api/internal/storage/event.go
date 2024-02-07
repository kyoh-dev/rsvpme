package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type Event struct {
	Uuid     string
	Title    string
	Descr    string       `db:"description"`
	Start    time.Time    `db:"start_datetime"`
	Finish   sql.NullTime `db:"finish_datetime"`
	Address  string
	Invitees []Invitee
}

func (d *Database) InsertEvent(e Event) (int, string, error) {
	var id int
	var uuid string
	q := `
		INSERT INTO event (title, description, start_datetime, finish_datetime, address) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, uuid
	`
	err := d.DB.QueryRowx(q, e.Title, e.Descr, e.Start, e.Finish, e.Address).Scan(&id, &uuid)
	if err != nil {
		return 0, "", fmt.Errorf("could not insert event: %w", err)
	}

	return id, uuid, nil
}
