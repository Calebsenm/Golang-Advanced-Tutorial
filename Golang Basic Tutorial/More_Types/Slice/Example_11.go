
package main 

import (
	"fmt"
)

func main(){
	numers1 := []int { 1 , 2 }
	numers2 := []int { 3 , 4 }

	numbers := append(numers1 , numers2...)

	fmt.Printf("numbers=%v\n", numbers)
    fmt.Printf("length=%d\n", len(numbers))
    fmt.Printf("capacity=%d\n", cap(numbers))
}