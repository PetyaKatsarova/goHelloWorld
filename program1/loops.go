package main

import (
	"fmt"
)

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println("tha val of i is:", i)
	}

	names := []string{"one", "two", "three", "four", "five"}

	for i := 0; i < len(names); i++ {
		fmt.Println("slice i is:", names[i])
	}

	for indx, val := range names {
		fmt.Printf("[%d] %s\n", indx, val)
	}

	for _, val := range names {
		fmt.Printf("the value is: %s\n", val)
	}
}