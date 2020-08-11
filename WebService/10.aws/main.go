package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1>Welcome</h1><p>To AWS hosted app</p>`)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}