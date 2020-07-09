package main

import (
	"fmt"
)

func main() {
	// #1
	a := 42
	fmt.Println(&a)

	// #2
	p1 := person {
		first: "Kinglon",
	}
	fmt.Println(p1)
	// changeMe(&p1)
	p1.changeMe()
	fmt.Println(p1)
}

func changeMe(p *person) {
	buf := ""
	
	for i := len(p.first) - 1; i > -1; i-- {
		buf = buf + string(p.first[i])
	}

	p.first = buf
}

func (p *person) changeMe() {
	buf := ""
	
	for i := len(p.first) - 1; i > -1; i-- {
		buf = buf + string(p.first[i])
	}

	p.first = buf
}

type person struct {
	first string
}