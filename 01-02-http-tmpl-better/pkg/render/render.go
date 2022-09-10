package render

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//create template cache

	//get reused template from cache

	//render the template

	tpl, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	custerror.ErrPrint(err)
	err = tpl.ExecuteTemplate(w, tmpl, nil)
	custerror.ErrPrint(err)
}

func createTemplateCache() (map[string]*template.Template, error) {

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
