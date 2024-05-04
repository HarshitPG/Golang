package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("x:",x)
		x++	
	}
	fmt.Print("\n")

	for i:=0;i<5;i++{
		fmt.Println("i",i)
	}
	fmt.Print("\n")

	char:= []string{"a","b","c","d"}
	for i:=0;i<len(char);i++{
		fmt.Println(char[i])
	}
	fmt.Print("\n")

	for index,value := range char{
		fmt.Printf("The value at index %v is %v \n",index,value)
	}
	fmt.Print("\n")

	for _,value := range char{
		fmt.Printf("The value is %v \n",value)
		value = "adding new char"// cant change new string to value bcuz its a shallow copy of the original string 
	}

	fmt.Println(char)
}