package main

import "fmt"

func main() {
	// #1
	fmt.Println(foo())
	_, b := bar()
	fmt.Print(b, " ")
	fmt.Println(bar())

	// #2
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(foo2(data...))
	fmt.Println(bar2(data))

	// #3
	defer defe()
	fmt.Println("defer 1")

	// #4
	p1 := person{
		first: "Jane",
		last: "Fonda",
		age: 18,
	}
	p1.speak()

	// #5
	sq := square {
		length: 2.5,
	}
	ci := circle {
		radius: 2.5,
	}
	info(sq)
	info(ci)

	// #6
	func(x string) {
		fmt.Println("from anonymous func with msg: ", x)
	}("message")

	//#7
	anon := func(sk int) int {
		return sk * sk
	}
	fmt.Println(anon(5))

	// #8
	sum := sum(4)
	fmt.Println(sum(4, 5))

	// #9
	fmt.Println(sumValue(foo2, data...))

	// #10
	inc := incrementor()
	fmt.Println(inc(), inc(), inc())
}

// #1
func foo() int {
	return 5
}
func bar() (int, string) {
	return 10, "abc"
}

// #2
func foo2(x ...int) int {
	res := 0
	for pos, val := range x {
		fmt.Print(pos)
		res += val
	}
	fmt.Println()
	return res
}
func bar2(y []int) int {
	res := 0
	for _, val := range y {
		res += val
	}
	return res
}

// #3
func defe() {
	defer func() {
		fmt.Println("defer 3")
	}()
	fmt.Println("defer 2")
}

// #4
type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Printf("Hello my name is %s %s, my age is %d\n", p.first, p.last, p.age)
}

// #5
type circle struct {
	radius float64
}
type square struct {
	length float64
}
func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}
func (s square) area() float64 {
	return s.length * s.length
}
type shape interface {
	area() float64
}
func info(s shape) {
	fmt.Println(s.area())
}

// #8
func sum(in int) func (int, int) int {
	return func (x int, y int) int {
		return (x + y) * in
	}
}

// #9
func sumValue(f func(...int) int, data ...int) []int {
	sum := f(data...)
	for pos, val := range data {
		data[pos] = val + sum
	}
	return data
}

// #10
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}