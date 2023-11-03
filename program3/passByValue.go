package main

import (
	"fmt"
)

func updateName(x string) {
	x = "tralala"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types: NON-POINTER VALUES -> strings, ints, bools, floats, arrays, structs: not changed in an array: need return value and to allocate it
	name := "tifa"
	updateName(name)
	fmt.Println(name) // tifa, not tralala

	// group B types: POINTER WRAPPER VALUES -> slices, maps, functions
	menu := map[string]float64 {
		"pie": 5.95,
		"ice cream": 3.99,
	}
	updateMenu(menu) // but here it did add coffee to the map!
	fmt.Println(menu)
}
