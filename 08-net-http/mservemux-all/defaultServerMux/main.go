package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

func c(res http.ResponseWriter, req *http.Request) {

	io.WriteString(res, "cat cat cattt")
}
func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	http.ListenAndServe("127.0.0.1:8080", nil)

}
