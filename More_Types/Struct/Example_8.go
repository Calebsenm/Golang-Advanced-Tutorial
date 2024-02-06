package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Employee struct {
    Name   string `json:"n"`
    Age    int    `json:"a"`
    Salary int    `json:"s"`
}

func main() {
    emp := Employee{Name: "Sam", Age: 31, Salary: 2000}
    //Converting to jsonn
    empJSON, err := json.MarshalIndent(emp, " ",  "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Println(string(empJSON))
}