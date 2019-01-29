package main

import (
	"go-examples/addressbook/contact"
	"go-examples/addressbook/storage"

	"github.com/go-kit/kit/log"
)

// AddressbookService is basic
type AddressbookService interface {
	LookFor(name string, lastName string) (contact.Contact, error)
}

type addressbookService struct {
	logger  log.Logger
	storage storage.Storage
}

func (srv addressbookService) LookFor(name string, lastName string) (contact.Contact, error) {
	cnt, err := srv.storage.LookFor(name, lastName)

	if err != nil {
		srv.logger.Log("msg", "Error fetching entry in storage", "err", err.Error())
		return contact.Contact{}, err
	}

	return cnt, nil
}
