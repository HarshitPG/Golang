package main

import "fmt"

func main() {
	num := 45

	fmt.Println(num<=50)
	fmt.Println(num>=50)
	fmt.Println(num==50)
	fmt.Println(num!=50)

	if num<30{
		fmt.Println("num less then 30")
	}else if num<40{
		fmt.Println("num less then 40")
	}else{
		fmt.Println("num not less then 40")
	}

	char:= []string{"a","b","c","d"}
	for index,value := range char{
		if index == 1{
			fmt.Println("cont. at index", index)
			continue
		}
		if index>2{
			fmt.Println("breakin at index", index)
			break
		}
		fmt.Printf("The value at index is %v is %v \n",index,value)
	}
}