package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type data struct {
	Method      string
	Submissions url.Values
	URL         *url.URL
	Header      http.Header
}

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	d := data{
		Method:      req.Method,
		Submissions: req.Form,
		URL:         req.URL,
		Header:      req.Header,
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", d)
	if err != nil {
		log.Fatalln(err)
	}

}
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog

	http.ListenAndServe("127.0.0.1:8080", d)
}
