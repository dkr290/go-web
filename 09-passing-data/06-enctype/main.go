package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
	Body       string
}

func init() {

	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {

	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	p := person{
		FirstName:  f,
		LastName:   l,
		Subscribed: s,
		Body:       body,
	}

	err := tpl.ExecuteTemplate(res, "index.gohtml", p)

	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}
