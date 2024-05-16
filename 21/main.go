package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return 4 * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printInfo(s shape){
	fmt.Printf("area of %T is: %0.2f \n",s,s.area())
	fmt.Printf("circumference of %T is: %0.2f \n",s,s.circumf())
}

func main(){
	shapes:= []shape{
		square{length: 2.5},
		circle{radius: 1.35},
		square{length: 5.5},
		circle{radius: 5},
	}

	for _,v := range shapes{
		printInfo(v)
		fmt.Println("")
	}
}
