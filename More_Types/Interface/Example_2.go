package main  

import "fmt"

type animal interface{
	breathe();
	walk();
}

type lion struct{
	age int 
}

func (l lion) walk(){
	fmt.Println("lion Walk");
}


func main(){

	var a animal
	a = lion{age: 10}
	a.breathe()
	a.walk()
	
}