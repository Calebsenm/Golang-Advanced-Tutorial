package main

import "fmt"

func main() {

	fmt.Println(add(2, 3, 5, 6, 6))
}

func add(numbers ...int) int {
	sum := 0
	
	for _, mum := range numbers {
		sum += mum
	}

	return sum
}
