package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input,err:=r.ReadString('\n')
	return strings.TrimSpace(input),err
}

func createBill() bill{
	reader:= bufio.NewReader(os.Stdin)

	// fmt.Print("Create new bill name:")
	// name,_ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name,_:=getInput("Create new bill name:",reader)

	b:= newBill(name)
	fmt.Println("Created a bill name:",b.name)

	return b

}

func promptOptions (b bill){
	reader:=bufio.NewReader(os.Stdin)
	opt,_:=getInput("Choose option (a- add item, s-save bill, t- add tip): ",reader)
	switch opt{
	case "a":
		name,_:=getInput("Item name:", reader)
		price,_:=getInput("Item price:", reader)
		p,err:=strconv.ParseFloat(price,64)
		if err != nil{
			fmt.Println("it must be a num")
			promptOptions(b)
		}
		b.updateItem(name,p)
		fmt.Println("item added -",name,p)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("selected s, bill saved", b.name)

	case "t":
		tip,_:=getInput("Enter Tip:", reader)
		t,err:=strconv.ParseFloat(tip,64)
		if err != nil{
			fmt.Println("it must be a num")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added -",t)
		promptOptions(b)
	default:
		fmt.Println("choose correct option")
		promptOptions(b)
	}
}

func main() {
	newbill:=createBill()
	promptOptions(newbill)
}