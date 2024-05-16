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
		items: map[string]float64{"pen": 1.99, "pencil": 0.5},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v  = $%v \n", k+":",v)
		total +=v
	}
	fs +=fmt.Sprintf("%-25v  = $%v \n", "Total:", total)
	return fs
}