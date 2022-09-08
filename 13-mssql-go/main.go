package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	db, err := sql.Open("mysql", "user:password@tcp(sql)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/vaficon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {

	_, err := io.WriteString(res, "Sucessfully completed")
	check(err)
}

func check(err error) {

	if err != nil {
		fmt.Println(err)
	}

}
