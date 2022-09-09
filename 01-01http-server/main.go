package main

import (
	"fmt"
	"net/http"
)

const hostPort = "127.0.0.1:8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("starting application on port", hostPort)
	http.ListenAndServe(hostPort, nil)

}

//this is the about page handler
func About(w http.ResponseWriter, req *http.Request) {
	sum := AddValues(3, 8)
	fmt.Fprintf(w, "This is the about page and sum is %v", sum)
}

//this is home page handler
func Home(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "This is the Home page")
	errPrint(err)

	//	fmt.Println("number of bytes written", n)
}

//add two integers and returns sum
func AddValues(x, y int) int {
	var sum int = x + y
	return sum

}

func errPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
