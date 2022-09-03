package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("cat.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/newcat/", newCat)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "<h2>foo run</h2>")

}

func cat(res http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(res, "cat.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func newCat(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}
