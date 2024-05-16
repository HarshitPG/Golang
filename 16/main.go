package main

import "fmt"

func main() {
	bill := newBill("John")
	bill.updateTip(10)
	bill.updateItem("Pen",5.99)
	bill.updateItem("Pencil",1.99)
	bill.updateItem("Note",7.5)
	bill.updateItem("Book",10)
	fmt.Println(bill.format())
}