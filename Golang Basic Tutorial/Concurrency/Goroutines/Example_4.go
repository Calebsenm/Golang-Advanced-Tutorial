
// Creating multiple goroutines 

package main 

import(
	"fmt"
	"time"
)

func executed(id int){
	fmt.Printf("Id: %d\n", id );
	
}

func main(){
	fmt.Println("Started");

	for i := 0; i < 10; i++{
		//time.Sleep(time.Microsecond * 3 );
		go executed(i)
	}

	time.Sleep(time.Second * 2 );
	fmt.Println("Finished");
}