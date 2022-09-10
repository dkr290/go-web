package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//get the template cache from app config

	//create template cache

	tc, err := CreateTemplateCache()
	custerror.FatalErr(err)

	//get reused template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}

	//render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all of the files page.gohtml

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	custerror.FatalErr(err)

	for _, page := range pages {
		//	first we need to get the actual name of the page using filepath.Base
		name := filepath.Base(page)
		// next, we need to actually create the template set parsing the file as well
		ts, err := template.New(name).ParseFiles(page)
		custerror.FatalErr(err)

		//	Next, the template set needs to know of any layouts we are using so it can parse correctly

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		custerror.FatalErr(err)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			custerror.FatalErr(err)
		}

		myCache[name] = ts
	}

	return myCache, nil

}
