package main

import (
	"net/url"
	"fmt"
	"net/http"
	"html/template"
	"strings"
	"log"
	"io"
)

type hotdog int
var temp *template.Template

func handleCat(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Cat-Key", "cat value")
	io.WriteString(res, "cat cat")
}

var fm template.FuncMap = template.FuncMap {
	"upper": testFunc,
}

func testFunc(s string) string {
	return strings.ToUpper(s)
}

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseGlob("./*.html"))
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Anything I can help with?")
	fmt.Println(r.URL, r.Body)
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Own-Key", "personal key here")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	data := struct {
		Method string
		URL *url.URL
		Submissions url.Values
		Header http.Header
		Host string
		ContentLength int64
	} {
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}

	temp.ExecuteTemplate(w, "formInput.html", data)
}

func main() {
	var d hotdog
	fmt.Println(d)

	// use DefaultServeMux with nil
	http.HandleFunc("/cat", handleCat)
	http.Handle("/cat/", http.HandlerFunc(handleCat))
	http.ListenAndServe(":8080", nil)
}