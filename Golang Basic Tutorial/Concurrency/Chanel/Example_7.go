
package main 

import (
	"fmt"
)

func main(){
	
	ch := make(chan int, 1 );
	ch <- 2 

	fmt.Println("Sending value to channel completed");

	val := <- ch 
	fmt.Printf("Receiving Value: %d", val )
}