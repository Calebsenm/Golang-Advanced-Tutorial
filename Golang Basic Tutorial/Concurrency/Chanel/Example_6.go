

package main 

import (
	"fmt"
)

func main(){

	ch := make(chan int );
	fmt.Println("Sending Value to Channel start");

	go send(ch);
	val := <- ch

	fmt.Printf("Receiving Value from channel finished: %d\n", val)

}

func send(ch chan int){
	ch <- 1
}