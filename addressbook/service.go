package main

// AddressbookService is basic
type AddressbookService interface {
	LookFor(name string, lastName string) (contact, error)
}

type addressbookService struct{}

type contact struct {
	name        string
	lastName    string
	address     string
	phoneNumber string
}

func (addressbookService) LookFor(name string, lastName string) (contact, error) {
	return contact{
			name:        "Name",
			lastName:    "Lastname",
			address:     "wonderful address",
			phoneNumber: "123123123"},
		nil
}
