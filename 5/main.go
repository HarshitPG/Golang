package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "hello world in go"

	fmt.Println(strings.Contains(text,"hello!"))
	fmt.Println(strings.ReplaceAll(text,"hello","hi"))
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.Index(text,"ll"))
	fmt.Println(strings.Split(text," "))

	//Original value is unchanged
	fmt.Println("original:",text)

	num:= []int{25,30,48,32,74,100,98}

	sort.Ints(num)
	fmt.Println(num)

	index:= sort.SearchInts(num,32)
	fmt.Println(index)

	char:= []string{"a","d","r","i","b","c"}

	sort.Strings(char)
	fmt.Println(char)
	fmt.Println(sort.SearchStrings(char,"b"))

}