package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) SetLastName(name string) {
	p.lastName = name
}

func (p *Person) SetFirstName(name string) {
	p.firstName = name
}

func main() {
	p := Person{}
	p.SetLastName("Doe")
	p.SetFirstName("John")
	fmt.Printf("Name: %s %s", p.GetFirstName(), p.GetLastName())
}
