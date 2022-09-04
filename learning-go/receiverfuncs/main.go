package main

import (
	"fmt"
	"log"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFn() string {

	return m.FirstName

}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("Myvar2 is set to", myVar2.FirstName)
	fmt.Println("Myvar is set to", myVar.FirstName)

	log.Println(myVar2.printFn())
	log.Println(myVar.printFn())
}
