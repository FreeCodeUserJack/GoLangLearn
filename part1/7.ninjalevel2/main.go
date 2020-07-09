package main

import(
	"fmt"
)

func main() {
	// #1
	a := 1024
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	// #2
	b := 4 <= 5
	c := 1 >= 4
	d := 1 == 1
	e := 2 != 1
	f := 55 > 2
	g := 66 < 5
	fmt.Println(b, c, d, e, f, g)

	// #3
	const (
		x = 6
		y string = "Hi"
		z = "Hello"
	)
	fmt.Println(x, y, z)

	// #4
	var val = 1024
	fmt.Printf("%d\t\t%b\t\t%#x\n", val, val, val)
	val <<= 1
	fmt.Printf("%d\t\t\t%b\t\t\t%#x\n", val, val, val)

	// #5
	var s string = `ewfawefaw
	wefawefawefaw			fewa;oij |_#%(%|_@ \n\t\n\t\n`
	fmt.Println(s)

	// #6
	const (
		i = 2017 + iota
		j uint32 = 2017 + iota
		k uint = 2017 + iota
		l int64 = 2017 + iota
	)
	fmt.Print(i, j, k, l, "\t\t")
	fmt.Printf("%T %T %T %T\n", i, j, k, l)

	// #7
	
}