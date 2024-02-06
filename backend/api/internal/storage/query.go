package storage

import "database/sql"

type Invitee struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     sql.NullString
	Rsvp      bool
}

func (d *Database) GetInvitee(id int) (Invitee, error) {
	inv := Invitee{}
	q := `
		SELECT first_name, last_name, email, rsvp
		FROM invitee
		WHERE id = ?
	`

	err := d.DB.Get(&inv, q)

	return inv, err
}

type Event struct {
	Title    string
	Descr    string
	Start    string
	End      string
	Address  string
	Invitees []Invitee
}
