
// Inner working of interface 

package main  

import (
	"fmt"
)

type animal interface{
	breathe()
	walk()
}

type lion struct{
	age int 
}


func (l lion) breathe(){
	fmt.Println("Lion Breathes")
}

func (l lion) walk(){
	fmt.Println("Lion Walk")
}

func main(){
	
	var a animal 
	a = lion{age: 10}
	fmt.Printf("Underying Type:  %T\n", a );
	fmt.Printf("Underlying Value: %v\n", a );
	
}