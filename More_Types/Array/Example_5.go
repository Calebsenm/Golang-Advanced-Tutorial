//Different ways of iterating an array

package main

import "fmt"

func main() {

	letters := [3]string{"a", "b", "c"}
	// Using for loop

	fmt.Println("Using for loop")
	len := len(letters)

	for i := 0; i < len; i++ {
		fmt.Println(letters[i])
	}

	// Using for-range operator
	fmt.Println("\n Using for range loop")

	for i, letter := range letters {
		fmt.Printf("%d %s \n ", i, letter)
	}
}
