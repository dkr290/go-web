package main

import "fmt"

func main() {
	fmt.Println("this is a test message")

	///variables creation

	var whatToSay string

	var i int

	whatToSay = "This is a test message"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	//short hand declaration
	whatWasSaid, theotherThingThatwasSaid := saySomething()

	fmt.Println(whatWasSaid, theotherThingThatwasSaid)

}

func saySomething() (string, string) {

	return "something", "else"

}
