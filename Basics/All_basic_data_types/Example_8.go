// Byte

package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main(){

	var r byte = 'a'

	// print Size 
	fmt.Printf("Size: %d \n", unsafe.Sizeof(r))

	// Print type 
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	// Print Character 
	fmt.Printf("Character: %c\n ", r )
	s := "abc"

	// This will the decimal value of bye
	fmt.Println([]byte(s))

}