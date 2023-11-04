package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Pip's bill")
	myBill.addItem("creme caramel", 4.42)
	myBill.addItem("green thi curry", 12)
	myBill.updateTip(10)
	fmt.Println(myBill.format())    
}