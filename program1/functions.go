package main

import (
	"fmt"
	"math"
)

func sayHi(n string) {
	fmt.Printf("hi %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("bye %v \n", n)
}


func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{ // return value is the last type typed
	return math.Pi * r * r
}

func main() {
	cycleNames([]string{"One", "Two", "Three"}, sayHi)
	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circl 2 is %0.3f", a1, a2)
}