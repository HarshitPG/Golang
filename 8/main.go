package main

import (
	"fmt"
	"math"
)

func sayHi(n string) {
	fmt.Printf("Hi %v \n",n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v\n",n)	
}

func cycleIPL(n []string, f func(string)) {
	for _,v := range n{
		f(v)
	}
	
}

func circleArea(r float64) float64{
	return math.Pi*r*r
	
}

func main() {
	sayHi("Harshit")
	sayBye("Harshit")

	cycleIPL([]string{"CSK","RCB","MI"},sayHi)
	cycleIPL([]string{"CSK","RCB","MI"},sayBye)

	a1:=circleArea(9.8)
	a2:=circleArea(10)
	fmt.Println("Area a1,a2:",a1,a2)

	fmt.Printf("circle a1 %0.3f, circle a1 %0.3f",a1,a2)
}