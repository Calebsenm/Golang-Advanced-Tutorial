package main

import "fmt"

func main(){

	sample := "afb"

	// with index and value
	fmt.Println("Both index and value")

	for i , letter := range sample{
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}

	// Only value
	fmt.Println("\n Only value")
	
	for _ , letter := range sample{
		fmt.Printf("Value: %s\n", string(letter))
	}

	// Only index 
	fmt.Println("\n Only Index ")
	for  i := range sample {
		fmt.Printf("Start Index: %d\n ", i )
	}
}