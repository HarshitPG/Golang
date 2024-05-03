package main

import "fmt"

func main() {

	//Array

	// var num [3]int =[3]int{1,2,3}
	var num = [3]int{1, 2, 3}
	fmt.Println(num,len(num))

	text:= [4]string{"one","two","three","four"}
	text[0]="changed"
	fmt.Println(text,len(text))

	//Slice

	var points = []int{1,3,5,2}
	points[2]=25
	points=append(points,101)
	fmt.Println(points,len(points))

	//slice ranges

	rangeone:= text[1:3]
	rangetwo:= text[1:]
	rangethree:= text[:3]
	fmt.Println(rangeone,rangetwo,rangethree)

	rangeone=append(rangeone, "change2")
	fmt.Println(rangeone)


}