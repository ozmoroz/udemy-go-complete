package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", contact: contactInfo{
		email:   "alex@gmail.com",
		zipCode: "94000",
	}}
	fmt.Printf("%+v", alex)
}
