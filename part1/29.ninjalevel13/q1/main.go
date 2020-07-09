package main

import (
	"fmt"
	"github.com/FreeCodeUserJack/GoLangLearn/29.ninjalevel13/q1/dog"
)

type canine struct {
	name string
	age int
}

func main() {
	fido := canine {
		name: "Fido",
		age: dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}