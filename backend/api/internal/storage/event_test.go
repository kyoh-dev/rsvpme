package storage

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertEvent(t *testing.T) {
	e := Event{
		Uuid:    uuid.NewString(),
		Title:   "Street Party",
		Descr:   "Party on Sesame Street",
		Start:   time.Now(),
		Finish:  sql.NullTime{Time: time.Now().Add(time.Hour * 10), Valid: true},
		Address: "1 Sesame Street",
	}
	_, err := testDb.InsertEvent(e)

	assert.NoError(t, err)
}

func setupReadEvent() *Event {
	var eventId int
	expEvent := Event{
		Uuid:    uuid.NewString(),
		Title:   "Bert's Open House",
		Descr:   "Bert's selling his house! Come by and view the package",
		Start:   time.Now(),
		Finish:  sql.NullTime{Time: time.Now().Add(time.Hour * 6), Valid: true},
		Address: "5 Sesame Street",
	}

	err := testDb.DB.QueryRowx(
		"INSERT INTO event (uuid, title, description, start_datetime, finish_datetime, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		expEvent.Uuid, expEvent.Title, expEvent.Descr, expEvent.Start, expEvent.Finish.Time, expEvent.Address,
	).Scan(&eventId)
	if err != nil {
		log.Fatalf("could not insert test event: %s", err)
	}
	testDb.DB.MustExec(
		"INSERT INTO invitee (first_name, last_name, email, rsvp, event_id) VALUES ($1, $2, $3, $4, $5)",
		"Elmo", "Sesame", "elmo.sesame@gmail.com", true, eventId,
	)

	return &expEvent
}

func TestReadEventByUuid(t *testing.T) {
	assert := assert.New(t)

	e := setupReadEvent()
	a, err := testDb.ReadEventByUuid(e.Uuid)

	assert.NoError(err)

	// TODO: Better testing for Event struct
	assert.Equal(e.Uuid, a.Uuid)
	assert.Equal(e.Title, a.Title)
	assert.Equal(e.Descr, a.Descr)
	assert.Equal(e.Address, a.Address)
}
