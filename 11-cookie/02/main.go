package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/bundle", bundle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {

	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")

}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNoContent)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE1:", c1)

	}

	c2, err := req.Cookie("general")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNoContent)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE2:", c2)

	}

	c3, err := req.Cookie("specific")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNoContent)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE3:", c3)

	}

}

func bundle(res http.ResponseWriter, req *http.Request) {

	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}
