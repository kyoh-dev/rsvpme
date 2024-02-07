package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Uuid     string
	Title    string
	Descr    string       `db:"description"`
	Start    time.Time    `db:"start_datetime"`
	Finish   sql.NullTime `db:"finish_datetime"`
	Address  string
	Invitees *[]Invitee
}

func (d *Database) InsertEvent(e Event) (int, string, error) {
	var id int
	uuid := uuid.New().String()
	q := `
		INSERT INTO event (uuid, title, description, start_datetime, finish_datetime, address) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id
	`
	err := d.DB.QueryRowx(q, uuid, e.Title, e.Descr, e.Start, e.Finish, e.Address).Scan(&id)
	if err != nil {
		return 0, "", fmt.Errorf("could not insert event: %w", err)
	}

	return id, uuid, nil
}

func (d *Database) ReadEventByUuid(uuid string) (*Event, error) {
	var e Event
	q := `
		SELECT 
			uuid, 
			title, 
			description, 
			start_datetime, 
			finish_datetime, 
			address,
			invitees
		FROM event_detail
		WHERE uuid = $1
	`
	err := d.DB.Get(&e, q, uuid)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
