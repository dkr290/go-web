package main

import (
	"log"
	"os"
	"text/template"
)

type items struct {
	Before   string
	MainDish string
	Drink    string
	Desert   string
	Count    int
}

type menu struct {
	MealTime string
	Menu     []items
}

func (i items) CountItems() int {

	return i.Count

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	type Menus []menu

	r := Menus{
		menu{
			MealTime: "Breakfast",
			Menu: []items{
				{
					Before:   "Salad",
					MainDish: "Beef roasted",
					Drink:    "Wine",
					Desert:   "Tarte fromage",
					Count:    1,
				},
			},
		},
		menu{
			MealTime: "Lunch",
			Menu: []items{
				{
					Before:   "Salad tomatoes",
					MainDish: "Pork with rice",
					Drink:    "Coca cola and Beer",
					Desert:   "Chocolate",
					Count:    1,
				},
			},
		},
		menu{
			MealTime: "Dinner",
			Menu: []items{
				{
					Before:   "Salad Concombre cheese",
					MainDish: "Chicken with rice",
					Drink:    "Coca cola and Beer",
					Desert:   "Chocolate",
					Count:    1,
				},
				{
					Before:   "Pasta Bolognese",
					MainDish: "From Italy delicious eating",
					Drink:    "Beer",
					Desert:   "Tarte",
					Count:    2,
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
