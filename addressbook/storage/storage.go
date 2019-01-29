package storage

import (
	"go-examples/addressbook/contact"
)

// Storage is where contacts are stored
type Storage interface {
	Save(cnt contact.Contact) (interface{}, error)
	LookFor(name string, lastname string) (contact.Contact, error)
}
