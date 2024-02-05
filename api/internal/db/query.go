package db

type Invitee struct {
	FirstName string
	LastName  string
	Email     string
	Rsvp      bool
}

func (d *Database) GetInvitee(uuid string) Invitee {
	return Invitee{}
}

type Event struct {
	Title    string
	Descr    string
	Start    string
	End      string
	Address  string
	Invitees []Invitee
}
