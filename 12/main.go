package main

import (
	"fmt"
)

func upadateName(a string) {
	a = "Johnson"
}

func upadateNameNew(b string) string{
	b = "Johnson"
	return b
}

func upadateShop(m map[string]float64){
	m["note"]=15.25
	m["Book"]=   1.10
}

func main() {
	// group A types => string, int, float, boolan, arrays, structs.
	// non-pointer values
	// store/update the value by creating a copy memory address

	name := "Jhon"

	upadateName(name)

	fmt.Println(name)

	//To update the memeory loaction return the value

	name=upadateNameNew(name)
	fmt.Println(name)


	// group B types => slices, map, functions
	// pointer wrapper values 
	// store/update the vulue in the same memory address
	shop := map[string]float64{
		"Pen":    1.99,
		"Pencil": 2.99,
		"Book":   10,
	}

	fmt.Println(shop)
	upadateShop(shop)
	fmt.Println(shop)
}