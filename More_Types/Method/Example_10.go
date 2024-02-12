
// Methods on Non-Struct Types 

package main 

import (
	"fmt"
	"math"
)

type myFLoat float64 

func (m myFLoat) ceil() float64{
	return math.Ceil(float64(m));
}

func main(){
	num := myFLoat(1.4);
	fmt.Println(num.ceil())
}
