package storage

import (
	"go-examples/addressbook/contact"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type redisStorage struct {
	network string
	host    string
	port    string
}

// NewRedisStorage creates a redis storage
func NewRedisStorage(network string, host string, port string) Storage {
	return redisStorage{
		network: network,
		host:    host,
		port:    port,
	}
}

func (r redisStorage) Save(cnt contact.Contact) (interface{}, error) {
	return nil, nil
}

func (r redisStorage) LookFor(name string, lastname string) (contact.Contact, error) {
	client, err := r.connect()
	if err != nil {
		return contact.Contact{}, err
	}
	defer client.Close()

	strJSON, err := redis.String(client.Do("GET", redisKey(name, lastname)))
	if err != nil {
		return contact.Contact{}, nil
	}

	return contact.NewContact(strJSON)
}

func (r redisStorage) connect() (redis.Conn, error) {
	return redis.Dial(r.network, r.host+":"+r.port)
}

func redisKey(name string, lastname string) string {
	return strings.ToLower(name + ":" + lastname)
}
