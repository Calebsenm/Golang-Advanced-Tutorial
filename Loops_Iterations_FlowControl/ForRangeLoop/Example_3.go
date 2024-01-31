
// For range loop with map 

package main

import "fmt"

func main(){

	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	//Interating over all keys and values
	fmt.Println("Both key and value");

	for k , v := range sample {
		fmt.Printf("Key :%s Value :%s \n " , k , v)
	}

	//Iterating over only keys 
	fmt.Println("\n Only Keys")
	
	for k := range sample {
		fmt.Printf("Key :%s \n", k )
	}

	// Iterating over only values 
	fmt.Println("\n Only Values")
	for _ , v := range sample {
		fmt.Printf("Value : %s\n ", v)
	}
	
}