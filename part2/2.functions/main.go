package main

import (
	"fmt"
	"strings"
)

func main() {
	//  variadic args b/c we must unfurl the slice with ...
	data := []string{"hello", "there"}
	a, b := greeting2(data...)
	fmt.Println(a, b)

	// func returning func
	getfunc := returnFunc()
	fmt.Println(getfunc("boo-ya!"))

	// closure
	clos1 := closure()
	clos2 := closure()
	fmt.Println(clos1(), clos2())
	fmt.Println(clos1(), clos2())

	// func taking in func
	var data2 []float64 = []float64{45.5, 54.5, 111.6, 909.5, 29292.3}
	fmt.Println(sumLessThanHundred(sumFloat, data2...))
}

// multi return variadic function with fmt.Spring and strings.Join()
func greeting2(s ...string) (string, string) {
	res := []string{}
	for _, x := range s {
		res = append(res, x)
	}
	return fmt.Sprint(strings.Join(res, " ")), "there"
}

// func returning func
func returnFunc() func(string) string {
	return func(ex string) string {
		return fmt.Sprint("got func ", ex)
	}
}

// closure
func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// func taking in func
func sumLessThanHundred(f func([]float64) float64, data ...float64) float64 {
	var toSum = make([]float64, len(data))
	for _, val := range(data) {
		if val < 100.0 {
			toSum = append(toSum, val)
		}
	}
	return f(toSum)
} 

func sumFloat(data []float64) float64 {
	res := 0.0
	for _, val := range data {
		res += val
	}
	return res
}