
// Non struct custon type implementing an interface 

package main  

import (
	"fmt"
)

type animal	interface{
	breathe()
	walk()
}

type cat string 

func (c cat) breathe() {
	fmt.Println("Cat breathes")
}

func (c cat ) walk(){
	fmt.Println("Cat Walk");
}

func main(){
	
	var	a animal 

	a = cat("Smokey")
	a.breathe()
	a.walk()
}