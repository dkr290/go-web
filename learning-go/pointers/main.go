package main

import (
	"log"
)

func main() {

	var myString string

	myString = "Green"

	log.Println("myString is set to", myString)

	changePointers(&myString)

	log.Println("after function call myString is changed to ", myString)
}

func changePointers(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
