
package main 

import "fmt"

func main(){

	go goRoutine()
	fmt.Println("Started")
	fmt.Println("Finished")
}

func goRoutine(){
	fmt.Println("Go Routine")
}