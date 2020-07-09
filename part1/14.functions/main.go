package main

import (
	"fmt"
)

type man struct {
	first string
}

type woman struct {
	last string
}

type human interface {
	speak() string
}

func (m man) speak() string {
	return "Me man"
}

func (w woman) speak() string {
	return "Me woman"
}

func greet(h human) {
	fmt.Println(h.speak())
	switch h.(type) {
	case man:
		fmt.Println("Man switch")
	case woman:
		fmt.Println("Woman switch")
	}
}

func main() {

	defer fmt.Println("this will be printed last")

	m1 := man {
		first: "James",
	}
	s, b := m1.foo(m1.first, 5)
	fmt.Println(s, b)

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...))

	w1 := woman {
		last: "Jane",
	}

	greet(m1)
	greet(w1)
	fmt.Printf("%T\n", m1)

	fun := func(x int) int {
		return x + 1
	}
	fmt.Printf("%T\n", fun)

	// returning a func
	f := bar()
	fmt.Println(f(4, 5))

	// callback
	fmt.Println(even(sum, 1, 2, 3, 4, 5))

	// recursion
	fmt.Println(fac(5))
}

func sum(x ...int) int {
	res := 0
	for _, val := range x {
		res += val
	}
	return res
}

// func callback, func as input
func even(f func(x ...int) int, vi ...int) int {
	slic := []int{}
	for _, val := range vi {
		if val % 2 == 0 {
			slic = append(slic, val)
		}
	}
	return f(slic...)
}

func (m man) foo(s1 string, i int) (string, bool) {
	for i < 10 {
		fmt.Println(s1)
		i += 2
	}
	b := false
	if i < 10 {
		b = true
	} else {
		b = false
	}
	return "done", b
}

// func returning func
func bar() func(int, int) int {
	f := func(a int, b int) int {
		return a + b
	}
	return f
}

// recursion
func fac(x int) int {
	if x == 2 {
		return 2
	}
	return fac(x-1) * x
}