// Integers

/*

Type	Size
int	    Platform Dependent
int8	8 bits/1 byte
int16	16 bits/2 byte
int32	32 bits/4 byte
int64	64 bits/8 byte

*/

package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main(){

	// This is a computed as  const uintSize = 32 << () (^uint(0) >> 32 & 1) // 32 or 64

	sizeOfIntBits := bits.UintSize
	fmt.Printf("%d bits \n" , sizeOfIntBits)

	var a int 

	fmt.Printf("%d buytes \n ", unsafe.Sizeof(a))
	fmt.Printf("a´s type is %s \n", reflect.TypeOf(a))

	b := 2 
	fmt.Printf("b ´s type is %s \n" , reflect.TypeOf(b))


}