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
		firstName: "Jom",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "jim@jim.com",
			zip:   1000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
