package main

import (
	"fmt"
	"net/http"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/handlers"
)

const hostPort = "127.0.0.1:8080"

func main() {

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)
	http.Handle("/favicon.cio", http.NotFoundHandler())
	fmt.Println("starting application on port", hostPort)
	http.ListenAndServe(hostPort, nil)

}
