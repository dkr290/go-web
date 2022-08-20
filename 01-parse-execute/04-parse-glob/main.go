package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {

	var err error
	tpl, err = template.ParseGlob("templates/*")

	//tpl = template.Must(template.ParseGlob("templates/*"))

	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
