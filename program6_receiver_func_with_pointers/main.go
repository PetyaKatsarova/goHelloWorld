package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Pip's bill")
	fmt.Println(myBill.format())    
}