package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

// format the bill: recieve the struct bill, so later we can do b.format()
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0;

	// list items
	for k, v := range b.items{
		fs += fmt.Sprintf("%-25v ... $%v\n", k+":", v) // all chars will take 25 chars
		total += v
	}
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

// update tip
func (b bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b bill) addItem(name string, price float64) {
	b.items[name] = price;
}