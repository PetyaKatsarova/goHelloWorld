package main

import (
	"fmt"
	"os"
)

// struct is like interface: just stating the structure of the object
type bill struct {
	name string
	items map[string]float64
	tip float64
}

// here create the object
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

// format the bill: recieve the struct bill, so later we can do b.format()
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0;

	// list items
	for k, v := range b.items{
		fs += fmt.Sprintf("%-25v ... $%v\n", k+":", v) // all chars will take 25 chars
		total += v
	}
	fs += fmt.Sprintf("%-25v ... $%v\n", "tip:", b.tip) // why doesnt update the tip???
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)
	return fs
}

// update tip RECEIVER FUNCTION!! WHEN RECEIVING A STRUCT
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip not needed, golang automatically deference it for us
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	(*b).items[name] = price;
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) // 0644 permissions, this returns an error
	if err != nil {
		panic(err) // stops the flow of the program
	}
	fmt.Println("bill was saved to file")
}