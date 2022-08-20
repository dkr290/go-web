package main

import (
	"html/template"
	"log"
	"os"
	"strings"
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

//"uc" is what the func is called in the template
//"ft" slices a string custom function returning first three chars

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s

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

	data := struct {
		Fruits    []fruit
		Transport []car
	}{
		frts,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
