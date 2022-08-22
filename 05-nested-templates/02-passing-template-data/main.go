package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "81 test string")
	if err != nil {
		log.Fatalln(err)
	}

}
