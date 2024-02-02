
package main

import (
	"fmt"
	"log"
	"os"
)



func main(){

	err := writeToTempFile("Some Text");
	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Println("Write to the file Succesful");

}

func writeToTempFile( test string ) error{

	file , err := os.Open("Temp.txt");

	if err != nil{
		return err
	}

	n , err := file.WriteString("Some Text");
	if err != nil {
		return err
	}

	fmt.Printf("Number of the bites Written :%d", n)
	file.Close()

	return nil
}
