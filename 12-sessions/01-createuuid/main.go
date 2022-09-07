package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}

		http.SetCookie(res, cookie)
	}

	fmt.Println(cookie)

}
