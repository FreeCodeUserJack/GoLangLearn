package main

import (
	// "fmt"
	"net/http"
	"io"
	// "os"
)

func dogPic(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="dog.png">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dogPic)
	http.ListenAndServe(":8080", nil)
}