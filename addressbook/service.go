package main

import (
	"github.com/go-kit/kit/log"
	"github.com/gomodule/redigo/redis"
	"os"
)

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

func redisDummyAction() {
	logger := log.NewLogfmtLogger(os.Stdout)

	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		logger.Log("msg", "Error conntecting to redis", "err", err.Error())
	}
	defer c.Close()

	ret, _ := c.Do("SET", "fleet", "truck1")
	logger.Log("set_ret", ret)

	ret, _ = c.Do("GET", "fleet")
	logger.Log("get_ret", ret)
}

func (addressbookService) LookFor(name string, lastName string) (contact, error) {

	redisDummyAction()

	return contact{
			name:        "Name",
			lastName:    "Lastname",
			address:     "wonderful address",
			phoneNumber: "123123123"},
		nil
}
