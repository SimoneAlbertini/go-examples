package contact

import "encoding/json"

// Contact represents the contact information for an entry in addressbook
type Contact struct {
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}

// ToJSON converts a Contact struct into a serialized json string
func (c Contact) ToJSON() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// NewContact creates a Contact struct from a json string
func NewContact(jsonStr string) (Contact, error) {
	bytes := []byte(jsonStr)
	var cnt Contact
	if err := json.Unmarshal(bytes, &cnt); err != nil {
		return Contact{}, err
	}

	return cnt, nil
}

//IsEmpty returns true if both last name and name are empty
func (c Contact) IsEmpty() bool {
	return c.Name == "" && c.LastName == ""
}
