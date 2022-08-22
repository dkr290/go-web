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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u1 := user{
		Name:  "James",
		Motto: "The first time movie is",
		Admin: false,
	}
	u2 := user{
		Name:  "Rohan",
		Motto: "This is a test account",
		Admin: true,
	}
	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
