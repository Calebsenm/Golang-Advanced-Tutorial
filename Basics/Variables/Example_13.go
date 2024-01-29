// // Scope of a Variable (Global)

package main 

import (
	"fmt"
)

var a  = "test"

func main(){
	testGlbal()
}

func testGlbal(){
	fmt.Println(a)
}