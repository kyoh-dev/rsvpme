package storage

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testDb Database
var elmoId int

func init() {
	testDb = Database{url: os.Getenv("DATABASE_URL")}
	if err := testDb.Connect(); err != nil {
		log.Fatalf("test database init failed: %s", err)
	}

	insertTestData(&testDb)
}

func insertTestData(db *Database) {
	testEvent := "INSERT INTO event (title, description, start_datetime, address) VALUES ($1, $2, $3, $4)"
	testInvitee := "INSERT INTO invitee (first_name, last_name, email, rsvp, event_id) VALUES ($1, $2, $3, $4, $5)"

	db.DB.MustExec(testEvent, "Elmo's House Party", "Party at Elmo's House!", time.Now(), "Sesame Street")

	var eventId int
	err := db.DB.Get(&eventId, "SELECT id FROM event WHERE title = $1", "Elmo's House Party")

	if err != nil {
		log.Fatalf("could not get eventId: %s", err)
	}

	db.DB.MustExec(testInvitee, "Elmo", "Sesame", "elmo.sesame@gmail.com", true, eventId)

	_ = db.DB.Get(&elmoId, "SELECT id FROM invitee WHERE first_name = $1", "Elmo")
}

func TestGetInvitee(t *testing.T) {
	actual, err := testDb.GetInvitee(elmoId)
	expected := Invitee{"Elmo", "Sesame", sql.NullString{String: "elmo.sesame@gmail.com", Valid: true}, true}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual, "unequal invitee structs")
}
