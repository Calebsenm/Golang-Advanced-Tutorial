
package main 

import (
	"fmt"
)

type base struct{
	color string
}

func (b *base) say(){
	fmt.Println("Hi from say function");
}

type child struct{
	base	// Enbedding 
	style	string 
}

func main(){

	base := base{color: "Red"}
	child := &child{
		base: base,
		style: "somesyle",
	}

	child.say();
	fmt.Println("The color is "+ child.color)
}