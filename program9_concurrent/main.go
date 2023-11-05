package main

import (
	"fmt"
	"time"
)

func main() {
	go f1("F1") // like async in js
	go f1("F2")
	fmt.Println("Sleeping for 5 sec")
	time.Sleep(5 * time.Second)
}

func f1(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v: i %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}
