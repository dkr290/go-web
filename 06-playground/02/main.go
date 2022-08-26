package main

import (
	"log"
	"os"
	"text/template"
)

type hotels struct {
	Name    string
	Address string
	City    string
	Zip     int
}

type regions struct {
	Region string
	Hotels []hotels
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	type Regions []regions

	h := Regions{
		regions{
			Region: "Southern",
			Hotels: []hotels{
				{Name: "California", Address: "South West Street 12", City: "Blabla", Zip: 12445},
				{Name: "Riviera", Address: "CEDEX 12", City: "Paris", Zip: 1244445},
				{Name: "Sofia Lux", Address: "Pentsho slavejkov 32", City: "Sofia", Zip: 12334343445},
			},
		},

		regions{
			Region: "Northern",
			Hotels: []hotels{
				{Name: "Plovdiv", Address: "South West Street 12", City: "Blabla", Zip: 98343},
				{Name: "Lux hotels 2", Address: "CEDEX 18", City: "Paris", Zip: 1244445},
				{Name: "Sofia Lux 44", Address: "Kiril 84", City: "Sofia", Zip: 12334343445},
			},
		},
		regions{
			Region: "Central",
			Hotels: []hotels{
				{Name: "Plovdiv", Address: "South West Street 12", City: "Blabla", Zip: 98343},
				{Name: "Lux hotels 2", Address: "CEDEX 18", City: "Paris", Zip: 1244445},
				{Name: "Sofia Lux 44", Address: "Kiril 84", City: "Sofia", Zip: 12334343445},
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}
