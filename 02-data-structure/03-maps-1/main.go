package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	fruits := map[int]string{
		1: "Apple",
		2: "Grapes",
		3: "Orange",
		4: "Prune",
		5: "Strawberry",
	}

	err := tpl.Execute(os.Stdout, fruits)
	if err != nil {
		log.Fatalln(err)
	}

}
