package main

import (
	"fmt"
)

var (
	//CompanyName represents the company name
	CompanyName      = "test"
	comapanyLocation = "somecity"
)

// Person struct
type Person struct {
	Name string
	age  int
}

// GetAge of person
func (p *Person) getName() string {
	return p.Name
}

type company struct{}

// GetPerson get the person object

func GetPerson() *Person {
	p := &Person{
		Name: "Test",
		age:  21,
	}

	fmt.Println("Model Package:")
	fmt.Println(p.Name)
	fmt.Println(p.age)

	return p
}

func main() {

}

func getCompanyName() string {
	return CompanyName
}
