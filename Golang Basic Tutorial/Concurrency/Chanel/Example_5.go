

package main 

import (
	"fmt"
	"time"
)

func main(){
	
	ch := make(chan  int )
	go send(ch)
	go recive(ch)

	time.Sleep(time.Second * 2)

}


func send(ch chan int ){
	time.Sleep(time.Second * 1);
	fmt.Println("Send data ")
	ch <- 1
}

func recive(ch chan int){
	val := <- ch 
	fmt.Printf("Data recived %d", val )
}
