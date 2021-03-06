package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jom",
		contactInfo: contactInfo{
			email: "jim@email.it",
			zip:   95488,
		},
	}

	jim.updateFirstName("jimmy")
	jim.print()
	jim.updateLastName("Goodman")
	jim.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) updateLastName(newLastName string) {
	p.lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
