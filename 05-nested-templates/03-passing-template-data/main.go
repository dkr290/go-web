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
	var r []int
	for i := 0; i < 50; i++ {
		r = append(r, i)
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}

}
