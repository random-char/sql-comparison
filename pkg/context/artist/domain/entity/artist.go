package entity

type Artist struct {
	id        int
	firstName string
	lastName  string
}

func NewArtist(id int, firstName string, lastName string) *Artist {
	return &Artist{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
	}
}

func (a *Artist) GetId() int {
	return a.id
}

func (a *Artist) GetFirstName() string {
	return a.firstName
}

func (a *Artist) GetLastName() string {
	return a.lastName
}

func (a *Artist) WithId(id int) *Artist {
	return &Artist{
		id:        id,
		firstName: a.firstName,
		lastName:  a.lastName,
	}
}
