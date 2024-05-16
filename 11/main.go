package main

import "fmt"

func main() {
	shop := map[string]float64{
		"Pen":    1.99,
		"Pencil": 2.99,
		"Book":   10,
	}

	fmt.Println(shop)

	//loops
	for k,v:=range shop{
		fmt.Println(k,"-",v)
	}

	fmt.Println(shop["Pen"])

	shop["Pen"]=5.15
	fmt.Println(shop)


}