package storage

import (
	"go-examples/addressbook/contact"
	"strings"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

const (
	testJsonString = "{\"name\":\"good\"," +
		"\"lastname\":\"boy\"," +
		"\"address\":\"such nice home\"," +
		"\"phonenumber\":\"099-111-bork\"}"
)

var (
	testContact = contact.Contact{
		Name:        "good",
		LastName:    "boy",
		Address:     "such nice home",
		PhoneNumber: "099-111-bork",
	}
)

func mockRedis() (s *miniredis.Miniredis, host string, port string) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	addr := strings.Split(s.Addr(), ":")
	host = addr[0]
	port = addr[1]
	return
}

func TestLookForExistingEntry(t *testing.T) {
	assert := assert.New(t)
	redisMock, host, port := mockRedis()
	defer redisMock.Close()

	redisMock.Set("foo:bar", testJsonString)

	storage := NewRedisStorage("tcp", host, port)
	contact, err := storage.LookFor("foo", "bar")

	assert.Nil(err)
	assert.Equal(testContact, contact)
}

func TestLookForNotExistingEntry(t *testing.T) {
	assert := assert.New(t)
	redisMock, host, port := mockRedis()
	defer redisMock.Close()

	storage := NewRedisStorage("tcp", host, port)
	contact, err := storage.LookFor("foo", "bar")

	assert.Nil(err)
	assert.True(contact.IsEmpty())
}

func TestSaveIsAStub(t *testing.T) {
	assert := assert.New(t)
	cnt, err := redisStorage{}.Save(testContact)
	assert.Nil(err)
	assert.Nil(cnt)
}
