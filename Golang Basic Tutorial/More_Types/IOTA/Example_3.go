
package main


import (
	"fmt"
)

type Size uint8 

const (
	small 		Size = 0
	medium 		Size = 1
	extraLarge 	Size = 2
	large 		Size = 3
)

func main(){

	fmt.Println(small);
	fmt.Println(medium);
	fmt.Println(large);
	fmt.Println(extraLarge);
}