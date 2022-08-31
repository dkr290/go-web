package main

import (
	"fmt"
	"net/http"
)

// interfaces how it works
type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	http.ListenAndServe("127.0.0.1:8080", d)

}
