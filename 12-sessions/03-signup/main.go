package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	Password string
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

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)

}

func signup(res http.ResponseWriter, req *http.Request) {

	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		// store user in dbUsers
		u := user{un, p, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "index.gohtml", u)

}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
