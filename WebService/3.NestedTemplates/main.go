package main

import (
	"fmt"
	"text/template"
	"os"
)

// Composition of data + method call
type person struct {
	Name string
}

type teacher struct {
	Person person
	Subject string
}

type student struct {
	Person person
	Grade string
}

type course struct {
	Prof teacher
	Title string
	Students []student
}

type school struct {
	Courses []course
	Students []student
	Profs []teacher
	Name string
}

func (p person) Speak() string {
	return fmt.Sprint("My name is: " + p.Name)
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	err := temp.ExecuteTemplate(os.Stdout, "parent.html", "Some")
	if err != nil {
		fmt.Println(err)
	}

	students := []student{
		student {
			Person: person{
				Name: "John",
			},
			Grade: "Freshman",
		},
		student {
			Person: person{
				Name: "Jack",
			},
			Grade: "Sophomore",
		},
		student {
			Person: person {
				Name: "Jane",
			},
			Grade: "Junior",
		},
	}

	profs := []teacher{
		teacher {
			Person:person{
				Name: "Kled",
			},
			Subject: "Math",
		},
		teacher {
			Person: person{
				Name: "Keith",
			},
			Subject: "Science",
		},
	}

	courses := []course{
		course {
			Prof: profs[0],
			Title: "Calculus",
			Students: students[:2],
		},
		course {
			Prof: profs[1],
			Title: "Biology",
			Students: students[1:],
		},
	}

	sch := school {
		Courses: courses,
		Students: students,
		Profs: profs,
		Name: "Vanderbilt",
	}
	err = temp.ExecuteTemplate(os.Stdout, "parent2.html", sch)
	if err != nil {
		fmt.Println(err)
	}
}