package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	lorem := person{
		firstName: "adsa",
		lastName:  "adad",
		contactInfo: contactInfo{
			email:   "faaaf",
			zipCode: 4566,
		},
	}

	lorem.updateName("cat")

	lorem.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
