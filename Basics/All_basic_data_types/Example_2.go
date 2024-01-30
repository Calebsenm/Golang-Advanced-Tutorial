// UnSigned

/*
Type	Size
uint	Platform Dependent
uint8	8 bits/1 byte
uint16	16 bits/2 byte
uint32	32 bits/4 byte
uint64	64 bits/8 byte

*/
package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main(){

	//This is computed as const uintSize = 32 << (^uuint(0) >> 32 & 1) // 32 or 64


	sizeOfuintInBits := bits.UintSize
    fmt.Printf("%d bits\n", sizeOfuintInBits)

    var a uint
    fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

}