package main

import "fmt"

func main(){
	//string
	var nameOne string = "Hi!"
	var nameTwo = "Hello"
	var nameThree string

	fmt.Println(nameOne,nameTwo,nameThree)

	nameOne="Hello"
	nameThree= "World"

	fmt.Println(nameOne,nameTwo,nameThree)

	nameFour := "Go"

	fmt.Println(nameFour)

	//ints

	var ageOne int = 20
	var ageTwo = 18
	ageThree :=50

	fmt.Println(ageOne,ageTwo,ageThree)

	//bits & memory 
	// var numOne int8 = -128
	// var numTwo uint16 = 256

	var floatOne float32 = 18.18
	var floatTwo float64 = 45656565977754.54154
	floatThree := 22.5

	fmt.Println(floatOne,floatTwo,floatThree)


}