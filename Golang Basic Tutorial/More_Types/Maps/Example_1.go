

package main 

import (
	"fmt"
)

func main(){
	
	// Declare 
	employeeSalary := map[string]int {}
	fmt.Println(employeeSalary)
	
	// Initialice using map lieteral 
	employeeSalary = map[string]int{
		"Joh": 1000,
		"Sam": 1200,
	}

	// Adding a key value 
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary);

	
}