package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"html/template"
)

type img struct {
	Width, Height int
	Title string
	Thumbnail thumbnail
	Animated bool
	IDs []int
}

type thumbnail struct {
	URL string
	Height, Wdith int
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("*"))
}

func main() {
	var data img

	rcvd := `{"Width":5450,"Height":3633,"Title":"Butterfly White Ling Insect","Thumbnail":
	{"Url":"https://cdn.pixabay.com/photo/2020/07/02/15/39/butterfly-5363370_1280.jpg","Height":125,"Width":100},
	"Animated":false,"IDS":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URL)

	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/ajax", ajax)
	http.HandleFunc("/foo", foo)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`hello`))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := struct {
		Name string
		Age int
	} {
		Name: "James",
		Age: 23,
	}
	js, err := json.Marshal(p1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(js)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := struct {
		Name string
		Age int
	} {
		Name: "Anna",
		Age: 25,
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ajax(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "ajax.html", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("foo called!")
	s := `Here is a sample response!`
	fmt.Fprintln(w, s)
}