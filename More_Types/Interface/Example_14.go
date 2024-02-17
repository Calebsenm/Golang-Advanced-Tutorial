
// Embedding Interfaces 

package main 
 
import "fmt"

type animal interface{
	breathe()
	walk()
}

type human interface{
	animal 
	speak()
}

type employee struct{
	name string 
}

func ( e employee) breathe(){
	fmt.Println("Employee Breathes")
}

func (e employee) walk(){
	fmt.Println("Employee walk ")
}

func (e employee) speak(){
	fmt.Println("Employee Speaks ")
}

func main(){
	var h human 
	h = employee{name: "John"}
	h.breathe()
	h.walk()
	h.speak()
}