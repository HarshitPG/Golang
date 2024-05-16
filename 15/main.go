package main

import "fmt"

func main() {
	bill := newBill("John")
	fmt.Println(bill.format())
}