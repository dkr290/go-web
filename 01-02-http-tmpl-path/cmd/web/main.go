package main

import (
	"fmt"
	"net/http"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/config"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/handlers"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/render"
)

const hostPort = "127.0.0.1:8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	custerror.FatalErrString("cannot create template cache", err)

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.Newtemplates(&app)

	http.Handle("/favicon.cio", http.NotFoundHandler())
	fmt.Println("starting application on port", hostPort)

	srv := http.Server{
		Addr:    hostPort,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	custerror.FatalErr(err)

}
