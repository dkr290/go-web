package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseFiles("tpl.gohtml"))
}

type score struct {
	Score1 int
	Score2 int
}

func main() {

	g1 := score{7, 9}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
