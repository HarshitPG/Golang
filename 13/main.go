package main

import "fmt"

func updateName(n string){
	n="Johnson"
}

func updateNameNew(n *string){
	*n="Johnson"
}

func main() {
	name := "John"

	// updateName(name)
	// fmt.Println("memory address of name is:", &name)

	m:=&name
	// fmt.Println("memory address:", m)
	// fmt.Println(name)

	updateNameNew(m)
	
	fmt.Println(name)
}