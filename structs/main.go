package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{
		email:   "alex@gmail.com",
		zipCode: "94000",
	}}
	alex.print()
	alex.updateName("Jimmy")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
