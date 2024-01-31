package main 

import "fmt"

func main(){

	switch char := "b"; char {
	case "a":
		fmt.Println("a");
	case "b":
		fmt.Println("b");
		break
	default: 
		fmt.Println("No matchign character");
	
	}
}