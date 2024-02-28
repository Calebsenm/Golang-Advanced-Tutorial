package model

import (
	"fmt"
)

// Test function

func Test() {
	// ESTRUCTURE IDENTIFIER
	p := &Person{
		Name: "test",
		age:  21,
	}

	fmt.Println(p)

	c := &company{}
	fmt.Println(c)

	// ESTRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Printlh(p.getName())

	// ESTRUCTURE'S FIELDS
	fmt.Pritnln(p.Name)
	fmt.Println(p.age)

	// FUNCTION
	person2 := GetPerson()
	fmt.Println(person2)

	companyName := getCompanyName()
	fmt.Println(companyName)

	// VARIBLES
	fmt.Println(companyLocation)
	fmt.Println(CompanyName)
}
