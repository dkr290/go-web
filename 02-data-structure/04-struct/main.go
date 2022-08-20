package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type fruits struct {
	Name  string
	Color string
	Descr string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	frt1 := fruits{
		Name:  "Orange",
		Color: "Yellow",
		Descr: "Oranges have vitamines",
	}

	err := tpl.Execute(os.Stdout, frt1)
	if err != nil {
		log.Fatalln(err)
	}

}
