package main

import (
	"fmt"
)

func updateName(x *string) { // * means the value of the pointer
	*x = "tralala"
}

func main() {
	// group A types: NON-POINTER VALUES -> strings, ints, bools, floats, arrays, structs: not changed in an array: need return value and to allocate it
	name := "tifa"
	m    := &name // & means the address of the var

	fmt.Println("string name: ", name) // tifa
	updateName(&name) // changes the value of the pointer
	fmt.Println(name) // tralala
	fmt.Println("value of m: ", m) // here dont need * dont know why: is m a pointer to an address? 

	updateName(m)
	fmt.Println("memory address of names is:", &name) // tifa, not tralala: show the address of the var!!

	fmt.Println("memory addres if  m is   : ", m)
	fmt.Println("value of memory addres is: ", *m) // deference: give value of the pointer
}
