// Typed constant

package main

import "fmt"

func main(){
	const a int32 = 8 
	
	var i1 int32 
	var i2 int32

	i1 = a 
	i2 = a 

	fmt.Println(i1 , "\n" , i2)
}