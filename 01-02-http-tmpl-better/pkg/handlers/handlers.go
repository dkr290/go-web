package handlers

import (
	"net/http"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/render"
)

// this is the about page handler
func About(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")

}

// this is home page handler
func Home(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}
