// Scope of a Variable (Local )

package main 

import "fmt"

func main(){
	
	var a = "test"
	fmt.Println(a)

	for i := 0 ; i < 3 ; i++ {
		fmt.Println(i)
	}

	fmt.Println(i)


}

func testLocal(){
	fmt.Println(a)
}