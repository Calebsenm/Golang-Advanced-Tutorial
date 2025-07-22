package main 

import(
    "fmt"
)

type Tv struct{
    isRunnig bool
}

func (t *Tv) on(){
    t.isRunnig = true;
    fmt.Println("Turning tv on"); 
}

func (t *Tv) off(){
    t.isRunnig = false;     
    fmt.Println("Turning tv off");
}
