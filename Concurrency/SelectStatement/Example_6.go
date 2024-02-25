
package main  

import (
	"fmt"
)

func main(){

	ch1 := make(chan  int )

	select {
	case mgs := <- ch1:
		fmt.Println(mgs);
	default:
		fmt.Println("Default statement executed");
	}
}
