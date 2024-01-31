// Higher order function 


package main

import (
	"fmt"
)

func main(){

	area := getAreaFunc()
	print(3 , 2 , area)

}

func print( x  ,  y int , area func( int , int ) int ){
	fmt.Printf("Area %d ", area( x , y ))
}

func getAreaFunc() func( int , int ) int {
	return func (x , y int ) int  {
		return x * y 
	}
}