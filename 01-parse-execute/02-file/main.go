package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	fs, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating the file", err)
	}
	defer fs.Close()

	err = tpl.Execute(fs, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
