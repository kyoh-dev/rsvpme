package storage

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertEvent(t *testing.T) {
	assert := assert.New(t)
	e := Event{
		Title:   "Street Party",
		Descr:   "Party on Sesame Street",
		Start:   time.Now(),
		Finish:  sql.NullTime{Time: time.Now().Add(1), Valid: true},
		Address: "1 Sesame Street",
	}
	_, _, err := testDb.InsertEvent(e)

	assert.NoError(err)
}

func TestReadEventByUuid(t *testing.T) {
	uuid := "ecd9d178-b8cd-4373-90fa-229c320340ca"
	assert := assert.New(t)

	q := "INSERT INTO event (uuid, title, description, start_datetime, address) VALUES ($1, $2, $3, $4)"
	testDb.DB.MustExec(q, "Open House", "Bert's open house inspection", time.Now(), "5 Sesame Street")

	e, err := testDb.ReadEventByUuid(uuid)
	assert.NoError(err)
	assert.Equal(uuid, e.Uuid)
}
