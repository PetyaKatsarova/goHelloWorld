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
