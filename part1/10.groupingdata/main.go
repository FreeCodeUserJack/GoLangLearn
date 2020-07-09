package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x [5]int
	var y [6]int
	fmt.Println(x)
	x[3] = 4
	fmt.Println(x)
	fmt.Printf("%T\t%T\n", x, y)

	// slice
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(a))
	fmt.Println(len(a))
	for p, e := range a {
		fmt.Printf("%d:%d ", p, e)
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d:%d ", i, a[i])
	}
	fmt.Println(a[1:], a[:1])
	a = append(a, 4, 5)
	fmt.Println(a)
	a = append(a[:2], a[3:]...)
	fmt.Println(a)

	mak := make([]int, 5, 10)
	fmt.Println(mak)

	msl := make([][]int, 4)
	fmt.Println(msl)
	for kk:=0; kk<4; kk++ {
		msl[kk] = make([]int, 3)
	}
	fmt.Println(msl)

	mapp := map[int] string{
		2: "fe",
		3: "fefe",
	}
	mapp[1] = "fewfa"
	v, ok := mapp[1]
	fmt.Println(v, ok)

	for k, v := range mapp {
		fmt.Println(k, v)
	}

	if _, ok := mapp[1]; ok {
		delete(mapp, 1)
	}
	fmt.Println(mapp)
}