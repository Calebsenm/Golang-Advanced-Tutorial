// A variable declared within an inner scope having the same name as variable declared in the outer scope will shadow the variable in the outer scope.

package main 

import (
	"fmt"
)

var a = 123

func main(){

	var a = 456
	fmt.Println(a)
}

