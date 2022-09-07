package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {

	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		sID := uuid.New()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: sID.String(),
		}

		http.SetCookie(res, cookie)
	}

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		fmt.Println("from un , ok", un)
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u

		fmt.Println(dbSessions)
		fmt.Println(dbUsers)
	}

	tpl.ExecuteTemplate(res, "index.gohtml", u)

}

func bar(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session-id")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
