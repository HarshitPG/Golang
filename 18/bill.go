package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v  = $%v \n", k+":",v)
		total +=v
	}
	fs +=fmt.Sprintf("%-25v  = $%v \n", "Tip:", b.tip)
	fs +=fmt.Sprintf("%-25v  = $%0.2f \n", "Total:", total+b.tip)
	return fs
}

// add items to the bill
func (b *bill) updateItem(name string, price float64){
	b.items[name]=price

}

//  add tip to the bill
func (b *bill) updateTip(tip float64){
	b.tip= tip
}