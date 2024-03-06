
package main 

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){

	//Declare a float64
	var a float64  = 2

	// size of a float 64 in bythes 
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	// Default is float64 when you do not specify a type 
	b := 2.3 
	fmt.Printf("b's type is %s\n", reflect.TypeOf(b))
	
}
