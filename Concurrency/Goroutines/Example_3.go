package main

import (
	"fmt"
	"time"
)

func main() {

	go start()


	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func start() {

	fmt.Println("In Goroutine")
	go start2()
}

func start2() {
	fmt.Println("In Goroutine2")
}
