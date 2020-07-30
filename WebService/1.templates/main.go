package main

import(
	"fmt"
	"os"
	"io"
	"strings"
	"text/template"
)

var temp2 *template.Template

func init() {
	temp2 = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	name := "Nate"
	// name := os.Args[1]
	fmt.Println(os.Args[0])

	tpl := fmt.Sprint(
`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` + name + `</h1>
</body>
</html>`)
	// fmt.Println(tpl)
	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))

	// parse a .html or .gohtml file
	temp, err := template.ParseFiles("index.html", "index2.html")
	if err != nil {
		fmt.Println(err)
	}
	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n")
	err = temp.ExecuteTemplate(os.Stdout, "index2.html", nil)
	fmt.Printf("\n\n")
	temp2.ExecuteTemplate(os.Stdout, "index.html", nil)
	fmt.Printf("\n\n")
	temp2.ExecuteTemplate(os.Stdout, "variable.html", "someVar")

	fmt.Printf("\n\n")
	sages := []string{"Jesu", "Muha", "Budd", "MLK", "Ghan"}
	temp2.ExecuteTemplate(os.Stdout, "variableSlice.html", sages)

	fmt.Printf("\n\n")
	mapp := map[string]int{"John": 1, "Jack": 2, "Jane":3}
	temp2.ExecuteTemplate(os.Stdout, "variableMap.html", mapp)

	type person struct {
		Name string
		Motto string
	}
	fmt.Printf("\n\n")
	buddha := person {
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}
	temp2.ExecuteTemplate(os.Stdout, "variableStruct.html", buddha)

	fmt.Printf("\n\n")
	persons := []person{
		person{"Buddha", "Buddhism"},
		person{"Jesus", "Christianity"},
		person{"MLK", "To have a dream"},
	}
	temp2.ExecuteTemplate(os.Stdout, "variableSliceStruct.html", persons)

	// not working
	type crowd struct {
		People []person
		Number int
		Desc string
	}
	fmt.Printf("\n\n")
	crowds := crowd {
		People: []person{
			person{"Jane", "Cool"},
			person{"Jack", "Awesome"},
		},
		Number: 3,
		Desc: "famous figures",
	}
	fmt.Println(crowds)
	temp2.ExecuteTemplate(os.Stdout, "variableStructSliceStruct.html", crowds)
}