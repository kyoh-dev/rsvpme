package storage

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
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
	t.Run("find record using uuid", func(t *testing.T) {
		uuid := uuid.New().String()
		assert := assert.New(t)

		q := "INSERT INTO event (uuid, title, description, start_datetime, address) VALUES ($1, $2, $3, $4, $5)"
		testDb.DB.MustExec(q, uuid, "Open House", "Bert's open house inspection", time.Now(), "5 Sesame Street")

		e, err := testDb.ReadEventByUuid(uuid)
		assert.NoError(err)
		assert.Equal(uuid, e.Uuid)
	})
}
