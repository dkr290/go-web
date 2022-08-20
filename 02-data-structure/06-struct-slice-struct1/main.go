package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type fruit struct {
	Name  string
	Color string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Fruits    []fruit
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := fruit{
		Name:  "Orange",
		Color: "Green",
	}

	g := fruit{
		Name:  "Grapes",
		Color: "White",
	}

	m := fruit{
		Name:  "Lemon",
		Color: "Yellow",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	frts := []fruit{b, g, m}
	cars := []car{f, c}

	data := items{
		Fruits:    frts,
		Transport: cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
