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
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jom",
		contact: contactInfo{
			email: "jim@email.it",
			zip:   95488,
		},
	}
	fmt.Printf("%+v", jim)
}
