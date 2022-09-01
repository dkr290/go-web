package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Test-Key", "is from test")
	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(res, "<h2>any code you want in this func</h2>")
}
func main() {
	var d hotdog

	http.ListenAndServe("127.0.0.1:8080", d)
}
