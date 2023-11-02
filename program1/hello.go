package main // every file belongs to a package, main is the default package for an executable program s opposed to a library
// When the Go build system sees a main package, it compiles the package into an executable binary, rather than a shared library.
// 

import (
	"fmt" // fmt package
	"errors"
)

func add(x int, y int) int {
	return x + y;
}

func divide(divident, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, errors.New("cant divide by zero")
	}
	return divident / divisor, nil
}


func main() {
		// struct
		var p Person
		p.Name = "John"
		p.Age = 42;
		p.SayHello()


	//  pointers
	a := 5
	b := &a // b is a pointer to a

	fmt.Println(a, b) // 5 0xc0000a0b8 , example memory locaiton
	fmt.Println(*b) // dereferencing b, to get the value of a

	result := add(1,2);
	fmt.Println("1 + 2 =", result)

	var msg string = "Hello, Go!"
	fmt.Println(msg)
	msg2 := "type inference" // no need to declare type var
	fmt.Println(msg2)

	for i := 0; i < 5; i++ {
		fmt.Println("Loop i:", i);
	}
}
