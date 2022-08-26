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
	Price    float64
}

type menu struct {
	MealTime string
	Menu     []items
}

type restaurant struct {
	Name string
	Menu []menu
}

func (i items) CountItems() int {

	return i.Count

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	type Restaurants []restaurant

	r := Restaurants{
		restaurant{
			Name: "2 Meals very good",
			Menu: []menu{
				{
					MealTime: "Breakfast",
					Menu: []items{
						{Before: "Salad", MainDish: "Roated Beef", Drink: "Wine", Desert: "tarte fromage", Count: 1, Price: 7.99},
						{Before: "Tsaziki", MainDish: "Beef with potatoes", Drink: "Wine or Beer", Desert: "tarte fromage", Count: 2, Price: 8.99},
					},
				},
				{
					MealTime: "Dinner",
					Menu: []items{
						{Before: "Salad", MainDish: "Roated Beef", Drink: "Wine", Desert: "tarte fromage", Count: 1, Price: 6.99},
						{Before: "Tsaziki", MainDish: "Pork with potatoes", Drink: "Wine or Beer", Desert: "Cheese", Count: 2, Price: 8.99},
					},
				},
			},
		},
		restaurant{
			Name: "The best Food in the North",
			Menu: []menu{
				{
					MealTime: "Lunch",
					Menu: []items{
						{Before: "Salad Tomatoes", MainDish: "Pork with rice", Drink: "Coca Cola or Beer", Desert: "Chocolate", Count: 1, Price: 4.87},
						{Before: "Salad Tomatoes", MainDish: "Chicken with rice", Drink: "Coca Cola or Beer or maybe Wine and Rakia", Desert: "Chocolate or tarte", Count: 2, Price: 7.47},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
