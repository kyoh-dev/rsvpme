package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx/types"
)

type Event struct {
	Uuid     string         `db:"uuid" json:"uuid"`
	Title    string         `db:"title" json:"title"`
	Descr    string         `db:"description" json:"description"`
	Start    time.Time      `db:"start_datetime" json:"startDatetime"`
	Finish   sql.NullTime   `db:"finish_datetime" json:"finishDatetime"`
	Address  string         `db:"address" json:"address"`
	Invitees types.JSONText `db:"invitees" json:"invitees"`
}

func (d *Database) InsertEvent(e Event) (int, error) {
	var id int
	q := `
		INSERT INTO event (uuid, title, description, start_datetime, finish_datetime, address) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id
	`
	err := d.DB.QueryRowx(q, e.Uuid, e.Title, e.Descr, e.Start, e.Finish, e.Address).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not insert event: %w", err)
	}

	return id, nil
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
		return nil, fmt.Errorf("could not read event from db: %s", err)
	}
	// TODO: Unmarshall invitees to slice of structs

	return &e, nil
}
