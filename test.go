package main

import (
	"fmt"
)

func main2() {
	fmt.Println("hi")
	foo()

	var a = 4
	b := 4

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("for looping", a, b)
		}
	}

	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("we exited")
}
