package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
)

func RenderTemplateNew(w http.ResponseWriter, tmpl string) {

	tpl, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	custerror.ErrPrint(err)
	err = tpl.ExecuteTemplate(w, tmpl, nil)
	custerror.ErrPrint(err)
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tpl *template.Template
	var err error

	//check if we have already the template in our cache

	_, ok := tc[t]
	if !ok {
		//need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		custerror.ErrPrint(err)

	} else {
		log.Println("Using cached tempalte")
	}
	tpl = tc[t]

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
	}
	fmt.Println(templates)
	//parse templates
	tpl, err := template.ParseFiles(templates...)
	custerror.ErrPrint(err)
	tc[t] = tpl
	fmt.Println(tc[t])

	return nil

}
