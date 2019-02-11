package contact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jsonString = "{\"name\":\"good\"," +
		"\"lastname\":\"boy\"," +
		"\"address\":\"such nice home\"," +
		"\"phonenumber\":\"099-111-bork\"}"
)

func TestCreateContactFromJson(t *testing.T) {
	assert := assert.New(t)

	contact, err := NewContact(jsonString)

	assert.Nil(err)
	assert.Equal("good", contact.Name)
}

func TestCreateContactReturnsErrorIfInvalidJson(t *testing.T) {
	assert := assert.New(t)

	jsonString := "{invalid json}"

	_, err := NewContact(jsonString)

	assert.Error(err)
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)
	assert.True(Contact{}.IsEmpty())
	assert.False(Contact{Name: "foo"}.IsEmpty())
	assert.False(Contact{LastName: "bar"}.IsEmpty())
	assert.False(Contact{Name: "foo", LastName: "bar"}.IsEmpty())

}

func TestContactToJson(t *testing.T) {
	assert := assert.New(t)

	json, _ := Contact{
		Name:        "good",
		LastName:    "boy",
		Address:     "such nice home",
		PhoneNumber: "099-111-bork",
	}.ToJSON()

	assert.Equal(jsonString, json)
}
