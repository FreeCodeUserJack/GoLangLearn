package main

import (
	"fmt"
	"text/template"
	"os"
	"strings"
	"time"
)

var temp *template.Template

var fm = template.FuncMap {
	"uc" : strings.ToUpper,
	"ft" : firstThree,
	"fdateMDY": monthDayYear,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseGlob("./*.html"))
}

func main() {
	// using funcs and PIPELINING
	data := "this is data"
	err := temp.ExecuteTemplate(os.Stdout, "func.html", data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	tpl := template.Must(template.New("tt").Parse("soething"))
	tpl.ExecuteTemplate(os.Stdout, "tt", nil)

	// formatting time
	fmt.Println()
	err = temp.ExecuteTemplate(os.Stdout, "time.html", time.Now())
	if err != nil {
		fmt.Println(err)
	}

	// global functions
	fmt.Println()

	type payload struct {
		Int1 int
		Int2 int
		Desc string
		Indices []string
	}

	datum := payload {
		Int1: 1,
		Int2: 2,
		Desc: "description",
		Indices: []string{"zero", "one", "two"},
	}
	err = temp.ExecuteTemplate(os.Stdout, "globalfuncs.html", datum)
	if err != nil {
		fmt.Println(err)
	}
}