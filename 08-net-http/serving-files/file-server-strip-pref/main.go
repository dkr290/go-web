package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", cat)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func cat(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/resources/toby.jpg">`)

}
