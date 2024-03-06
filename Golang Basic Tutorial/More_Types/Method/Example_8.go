
package main 

import(
	"fmt"
)

// Person Struct 
type Person struct{
	Name	string 
	age 	int 
}

// GetAge of Person 
func (p *Person) GetAge() int {
	return p.age
}

func (p *Person)getName() string{
	return p.Name
}

type company struct{

}

func main(){

	p := &Person{
		Name: "test",
		age: 21,
	}

	fmt.Println(p);

	c := &company{}
	fmt.Println(c);

	//STRUCTURE´S FIELDS
	fmt.Println(p.Name);
	fmt.Println(p.age);

	// STRUCTURE´S METHOD 
	fmt.Println(p.GetAge());
	fmt.Println(p.getName())
}