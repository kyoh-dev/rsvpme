package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInvitee(t *testing.T) {
	db := Database{}

	actual := db.GetInvitee("abc123")
	expected := Invitee{"Elmo", "Sesame", "elmo.sesame@gmail.com", true}

	assert.Equal(t, expected, actual, "Unequal invitee structs")
}
