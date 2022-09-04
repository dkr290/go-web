package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))

}
func main() {

	http.HandleFunc("/", getimages)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)

	http.ListenAndServe("127.0.0.1:8080", nil)

}

func getimages(res http.ResponseWriter, req *http.Request) {

	tpl.Execute(res, nil)

}
