package main

import (
	"fmt"
	"log"
)

var s = "seven"

func main() {

	var s2 = "six"

	fmt.Println("s is ", s)
	fmt.Println("s2 is an", s2)

	b1, b2 := saySomething("xxx")

	log.Println("s from the saySomthing is", b1, "and second value is", b2)

}

func saySomething(s string) (string, string) {

	return s, "World"

}
