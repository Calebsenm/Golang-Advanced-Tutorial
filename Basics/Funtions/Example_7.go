// IIF or immediately invoked function

package main

import "fmt"

func main(){


	area := func () int{
		return 2 + 2
	}()

	fmt.Println(area)
}