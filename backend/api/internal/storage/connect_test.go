package storage

import (
	"log"
	"os"
	"testing"
)

var testDb Database

func setup() {
	if err := testDb.Init(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatalf("test db init failed: %s", err)
	}
}

func shutdown() {
	q := `
		DELETE FROM event;
		DELETE FROM invitee;
	`
	testDb.DB.Exec(q)
}

func TestMain(m *testing.M) {
	setup()
	exec := m.Run()
	shutdown()
	os.Exit(exec)
}
