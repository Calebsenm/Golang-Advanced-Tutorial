
package main  

import (
	"fmt"
)

type employee struct{
	name 	string 
	age 	int 
	salary 	int 
}

func main(){
	emp := employee{name: "Sam" , age: 31 , salary: 2000 }

	// Accecing a strunct field 
	n := emp.name
	fmt.Printf("Current name is : %s \n" , n )

	// Assing a new value 
	emp.name = "John"

	fmt.Printf("New name is: %s\n ", emp.name )
}