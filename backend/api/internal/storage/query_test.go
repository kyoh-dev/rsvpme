package storage

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDb Database

func init() {
	if err := testDb.Connect(); err != nil {
		log.Fatalf("test database init failed: %s", err)
	}
}

func TestGetInvitee(t *testing.T) {
	actual, err := testDb.GetInvitee(123)
	expected := Invitee{"Elmo", "Sesame", sql.NullString{String: "elmo.sesame@gmail.com"}, true}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual, "unequal invitee structs")
}
