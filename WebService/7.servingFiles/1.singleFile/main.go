package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Println("inside dog handle func")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<h1>Dog</h1><img src="/dog.png">`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("dog.png")
	if err != nil {
		http.Error(res, "file not found", 404)
	}
	defer f.Close()

	// #1 use io.Copy
	// io.Copy(res, f)

	// #2 use http.ServeFile(res, req, "dog.png") // only need this 1 line in this func

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
	}
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}

func main() {
	// dog image from wikipedia, this is nonprofit use, for learning only
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.png", dogPic)
	http.ListenAndServe(":8080", nil)
}