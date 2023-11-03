package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64 {
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffe pudding": 3.55, // u need the , here too!
	}

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string {
		123456: "mario",
		234567: "luigi",
		345678: "PK",
	}
	fmt.Println(phonebook[123456])
}
