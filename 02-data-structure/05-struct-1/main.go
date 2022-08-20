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
	frt2 := fruits{
		Name:  "Apple",
		Color: "Green",
		Descr: "Apples are good for health",
	}

	frt3 := fruits{
		Name:  "Grapes",
		Color: "Black or White and Green",
		Descr: "You can make wine",
	}

	all_fruits := []fruits{frt1, frt2, frt3}

	err := tpl.Execute(os.Stdout, all_fruits)
	if err != nil {
		log.Fatalln(err)
	}

}
