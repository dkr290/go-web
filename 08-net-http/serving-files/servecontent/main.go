package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	http.ListenAndServe("127.0.0.1:8080", nil)

}

func dog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(res, `<img src="toby.jpg">`)

}

func dogPic(res http.ResponseWriter, req *http.Request) {

	f, err := os.Open("toby.jpg")
	if err != nil {

		http.Error(res, "File not found", 404)
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}
