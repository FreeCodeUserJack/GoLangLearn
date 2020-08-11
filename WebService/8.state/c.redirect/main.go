package main

import (
	"net/http"
	"fmt"
	"html/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("./*.html"))
}

func foo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.ExecuteTemplate(res, "form.html", nil)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.FormValue("name"), req.Method)

	// res.Header().Set("Location", "/")
	// res.WriteHeader(http.StatusSeeOther)
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}