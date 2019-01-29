package main

import (
	"fmt"
	"go-examples/addressbook/contact"
	"time"

	"github.com/go-kit/kit/log"
)

// It implements the StringService interface!
type loggingMiddleware struct {
	logger log.Logger
	next   AddressbookService
}

func (mw loggingMiddleware) LookFor(name string, lastname string) (output contact.Contact, err error) {
	defer func(begin time.Time) {
		contactJSON, _ := output.ToJSON()
		mw.logger.Log(
			"method", "lookfor",
			"input", fmt.Sprintf("name: %s, lastname: %s", name, lastname),
			"output", contactJSON,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.LookFor(name, lastname)
	return
}
