package main

import "fmt"

func main() {

	printType("Test_String ")
	printType(2)
}

func printType(t interface{}) {

	switch v := t.(type) {
	case string:
		fmt.Println("Type String ");
	case int:
		fmt.Println("Type: int ");
	default:
		fmt.Printf("Unknown Type %T", v )
	
	}
}