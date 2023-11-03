package main // every file belongs to a package, main is the default package for an executable program s opposed to a library
// When the Go build system sees a main package, it compiles the package into an executable binary, rather than a shared library.
//

import (
	"errors"
	"fmt" // fmt package
	"sort"
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

	// from manyReturnVals.go
	fn1, sn1 := getInitials("lala kuku")
	fmt.Println(fn1, sn1)
	fn2, sn2 := getInitials("petka motoretka")
	fmt.Println(fn2, sn2)
	fn3, sn3 := getInitials("petka")
	fmt.Println(fn3, sn3)

	// from functions.go
	cycleNames([]string{"One", "Two", "Three"}, sayHi)
	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circl 2 is %0.3f", a1, a2)

	// from loops.go
	loops()

		// slices use arrays under the hood: can append to a slice
		var scores = []int{100, 50, 60}
		fmt.Println(scores)
		scores = append(scores, 85)
		fmt.Println(scores)

		// arrays have always fixed length! cant append to an array
		// var ages [3]int = [3]int{20, 25, 30}
		var ages = [3]int{20, 25, 30}
		fmt.Print(ages, len(ages), "\n")
		names := [4]string{"me", "you", "Lala", "Kuku"}
		fmt.Println("names and names.length", names, len(names))
		// yes but cant sort array, only slice!! so: 
		names2 := []string{"me", "you", "Lala", "Kuku"}
		sort.Strings(names2)
		fmt.Println("sorted names:", names2)
		fmt.Println(sort.SearchStrings(names2, "Lala"))

		// slice ranges
		rangeOne := names[1:3]
		rangeTwo := names[2:] // till the end including last el
		rangeThree := names[:3] // start from 0 till position 3 excluding
		fmt.Println("rangeOne: ", rangeOne)
		fmt.Println("rangeTwo: ", rangeTwo)
		fmt.Println("rangeThree:", rangeThree)
		rangeOne = append(rangeOne, "appended")
		fmt.Println("rangeOne appended: ", rangeOne)
		// struct
		// var p Person
		// p.Name = "John"
		// p.Age = 42;
		// p.SayHello()


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
