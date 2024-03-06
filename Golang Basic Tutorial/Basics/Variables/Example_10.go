// Variable Expression - While the declaration can also be assigned an expression or a funcional call. see the below example.

package main

import (
	"fmt"
	"math"
)

func main(){

	a := 5 + 3 
	b := math.Max(4 , 5 )

	fmt.Println(a)
	fmt.Println(b)

	
}