package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	io.WriteString(res, "cat cat cattt")
}
func main() {

	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe("127.0.0.1:8080", mux)

}
