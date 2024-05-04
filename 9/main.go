package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	text := strings.Split(s," ")

	var letter []string
	for _,v := range text{
		letter = append(letter,v[:1])
	}
	if len(letter)>1 {
		return letter[0],letter[1]
	}
	return letter[0],"_"
}
func main() {
	s1,s2 := getInitials("pen pencil")
	fmt.Println(s1,s2)
	c1,c2:= getInitials("hi")
	fmt.Println(c1,c2)
}