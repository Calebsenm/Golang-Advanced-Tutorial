// Acces underlying variable of interface
// type assertion

package main

import (
	"fmt"
)

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion Breathes")
}

func (l lion) walk() {
	fmt.Println("Lion Walk")
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog Walks")
}

func main() {

	var a animal
	a = lion{age: 10}

	print(a)
}

func print(a animal) {
	l := a.(lion)
	fmt.Printf("Age: %d\n", l.age)

	//d := a.(dog)
	//fmt.Printf("Age: %d\n", d.age)
}
