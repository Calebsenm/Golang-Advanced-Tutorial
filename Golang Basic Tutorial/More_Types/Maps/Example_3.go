

// Using make 
package main

import (
	"fmt"
)

func main(){
	// Declare 

	employeeSalary := make( map[string] int )

	// Adding a key value 
	employeeSalary["Lucas"] = 2000 
	fmt.Println( employeeSalary );

}