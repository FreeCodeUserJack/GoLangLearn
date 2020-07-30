package main

import (
	"fmt"
	"runtime"
)

func main() {
	const (
		n1 = 1 << iota
		n2
		n3
	)
	fmt.Println(n1, n2, n3)
	fmt.Println(runtime.GOARCH)

	const (
		a =       99999999999999999999999999999999999
		// b int64 = 99999999999999999999999999999999999 // will overflow but untyped const won't (256 bits)
	)

	var s1 []string = make([]string, 5, 8)
	s1[0] = "A"
	s1[1] = "B"
	s1[2] = "C"
	s1[3] = "D"
	s1[4] = "E"

	s2 := s1[:2:2] // set capacity
	fmt.Printf("%v %v %p %p", s1, s2, &s1, &s2)

	s2 = append(s2, "Z")
	fmt.Printf("%v %v %p %p", s1, s2, &s1, &s2)
}