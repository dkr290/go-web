package render

import (
	"net/http"
	"text/template"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tpl, _ := template.ParseFiles("./templates/" + tmpl)
	err := tpl.Execute(w, nil)
	custerror.ErrPrint(err)
}
