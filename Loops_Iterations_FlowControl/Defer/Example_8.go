
// Defer and panic 

package main 

import (
	"fmt"
)

func main(){

	defer fmt.Println("Defer in main");

	panic("Panic with Defer");
}