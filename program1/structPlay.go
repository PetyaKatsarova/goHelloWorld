package main

import "fmt"


//a typed collection of fields, useful for grouping data together
type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hi my name is", p.Name)
}

