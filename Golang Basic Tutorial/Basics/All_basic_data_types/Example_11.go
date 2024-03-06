// Booleans 

package main 

import "fmt"

func main(){

	// default value be false it not initialized
	
	var a bool
	fmt.Printf("a's value is %t\n", a)

	// And operation on one true and other false 

	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("Ouput of AND operation on one true and other false %t\n", andOperation)
    
	//Or Oeration on one true and other false 
	orOperation := 1 < 2 || 1 > 3 

	fmt.Println("Ouput of OR operation on one true and other false: %t\n", orOperation)

	// Negation Operation on false value 

	negationOperation := !(1 > 2 )
	fmt.Printf("Ouput of NEGATION operation on false value: %t\n", negationOperation)

}