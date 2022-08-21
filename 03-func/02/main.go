package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}

func monthdayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"dateMDY": monthdayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
