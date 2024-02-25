
package main  

import "fmt"

func main(){
	ch := make(chan int, 3 );
	fmt.Println("Capacity: %d\n", cap(ch));

}