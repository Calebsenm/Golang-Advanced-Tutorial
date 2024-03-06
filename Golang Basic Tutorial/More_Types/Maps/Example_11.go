
// Iterating Over a map 
package	 main 

import (
	"fmt"
)

func main(){

	sample := map[string] string{
		"a": "x",
		"b": "y",
	}

	for k , v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v);
	}

}