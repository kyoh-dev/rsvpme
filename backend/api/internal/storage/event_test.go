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
