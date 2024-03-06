// Scheduling of the goroutines

package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.NumCPU())
}