package main

import "fmt"

func main(){
	age:=19
	name:="Harshit"

	//Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")

	//Println
	fmt.Println("Hello")
	fmt.Println("Im ",name ,",Im ",age,"old.")

	//Printf (formatted string) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n",age, name)
	fmt.Printf("my age is %q and my name is %q \n", age , name)
	fmt.Printf("name type is %T \n", name)
	fmt.Printf("you scored %f points \n",25.555 )
	fmt.Printf("you scored %0.1f points \n", 34.456)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("Hi Im %v", name)
	fmt.Println("the saved string is:", str)


}