
// for-range loop for array/slice

package main 

import "fmt"

func main(){

	letters := []string{ "a", "b", "c", "d", "e"}

	// With index and value 

	fmt.Println("Both index value");

	for i , letter := range letters{
		fmt.Printf("index: %d Value %s \n", i , letter)
	}

	// Only value
	fmt.Println("\nOnly Index")

	for i := range letters {
		fmt.Printf("Idex: %d\n ", i)
	}

	// without index and value . just print array values
	fmt.Println("\nWithout Idex and Vale");

	i := 0

	for range letters{
		fmt.Printf("Index %d Value: %s\n", i , letters[i])
		i++
	}

}