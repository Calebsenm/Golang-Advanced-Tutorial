// Continue statement in for loop

package main

import "fmt"

func main(){

	for  i  := 1 ; i < 10 ; i++ {
		
		if i%3 == 0{
			continue
		}
		fmt.Println(i)
	}

	
}
