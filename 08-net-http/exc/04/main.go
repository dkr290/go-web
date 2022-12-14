package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))

}

func main() {

	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func dogs(res http.ResponseWriter, req *http.Request) {

	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
