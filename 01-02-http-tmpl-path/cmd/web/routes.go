package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/config"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
