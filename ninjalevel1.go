package main

import (
	"fmt"
)

// #2 / 3
var a int = 42 
var b string = "James Bond"
var c bool = true

// #3 illegal to do outside func {}
// a = 42
// b = "James Bond"
// c = true

// #5
type mytype int
var e mytype

func main() {
	// main2()
	// #1
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// #2
	fmt.Println(a, b, c) // zero-values, string is ""

	// #3
	s := fmt.Sprintf("%d\t%s\t%v", a, b, c)
	fmt.Println(s)

	// #4
	type ninja int
	var d ninja
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	var b int
	b = int(d)
	fmt.Println(d, b)

	// #5
	e = 5
	var f int
	f = int(e)
	g := int(e)
	fmt.Println(e, f, g)
}