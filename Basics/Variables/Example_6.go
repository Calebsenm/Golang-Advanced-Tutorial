// Variable Declaration with no type or Type Inference

package main 

import (
	"fmt"
)

func main(){

	var t = 123 		// Type Infered will be int 
	var u = "Circle"    // Type Infered will be string 
	var v = 5.6         // Type Infered will float64
	var w = true        // Type Infered will bool 
	var x = 'a'         // Type Infered will rune 
	var y = 3 + 5i      // Type Infered will complex128
	var z = sample{name: "Test"}  //Type Inferred will be main.Sample


	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Printf("Type: %T Value: %v\n", u, u)
    fmt.Printf("Type: %T Value: %v\n", v, v)
    fmt.Printf("Type: %T Value: %v\n", w, w)
    fmt.Printf("Type: %T Value: %v\n", x, x)
    fmt.Printf("Type: %T Value: %v\n", y, y)
    fmt.Printf("Type: %T Value: %v\n", z, z)

}

type sample struct{
	name string
}
