package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"html/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("./*.gohtml"))
}

func foo (res http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
		defer f.Close()

		fmt.Println("\nfile: ", f, "\nheader: ", h, "\nerror: ", err)

		bs, err := ioutil.ReadAll(f) // bytestream return and error
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// store the file

		// unnecessary as the . for curr dir works
		// dir, err := os.Getwd()
		// fmt.Println(err, dir)

		// dst, err := os.Create(filepath.Join(dir, "\\", h.Filename))
		dst, err := os.Create(filepath.Join("./", h.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.ExecuteTemplate(res, "index.gohtml", s)
}