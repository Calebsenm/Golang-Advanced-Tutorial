// Floats

/*
Type	Size
float32	32 bits or 4 bytes
float64	64 bits or 8 bytes

*/

package main 

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){

	//Declare a float32
	var a float32 = 2

	// Size of flat32 in bytges 
	fmt.Printf("%d Bytes \n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n ", reflect.TypeOf(a))


}