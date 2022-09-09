package render

import (
	"net/http"
	"text/template"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tpl, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	custerror.ErrPrint(err)
	err = tpl.ExecuteTemplate(w, tmpl, nil)
	custerror.ErrPrint(err)
}
