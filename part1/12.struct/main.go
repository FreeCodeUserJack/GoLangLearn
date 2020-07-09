package main

import "fmt"

func main() {
	type person struct {
		first string
		last string
	}

	type secretA struct {
		person
		ltk bool
	}

	p := person{first:"jack", last:"daniels"}
	fmt.Println(p)

	sa1 := secretA {
		person: person {
			first: "Hi",
			last: "Jack",
		},
		ltk: true,
	}
	fmt.Println(sa1)
}