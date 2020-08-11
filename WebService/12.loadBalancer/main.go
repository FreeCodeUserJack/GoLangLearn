package main

import(
	"fmt"
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://instance_public_ip/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
	}

	bs := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(bs)
	if err != nil {
		fmt.Println(err)
	}

	resp.Body.Close()

	io.WriteString(w, string(bs))
}