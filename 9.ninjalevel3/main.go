package main

import (
	"fmt"
)

func main() {

	// #1
	for i := 1; i < 10000; i += 1000 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// #2
	for i := 65; i <= 91; i ++ {
		fmt.Println(i-64)
		j := 0
		for {
			if j == 3 {
				break
			}
			fmt.Printf("\t%#U ", i)
			j++
		}
		fmt.Println()
	}

	// #3
	k := 1994
	for k < 2021{
		fmt.Printf("%d ", k)
		k++
	}
	fmt.Println()

	// #4
	l := 1994
	for {
		if l == 2021 {
			break
		}
		fmt.Printf("%d ", l)
		l++
	}
	fmt.Println()

	// #5
	for m := 10; m < 101; m++ {
		fmt.Printf("%d ", m % 4)
	}
	fmt.Println()

	// # 6
	if true {
		fmt.Println("true")
	}

	// #7
	if 1 < 0 {
		fmt.Println("<1")
	} else if 1 == 1 {
		fmt.Println("1")
	} else {
		fmt.Println(">1")
	}

	// #8
	switch {
	case true:
		fmt.Println("true")
		fallthrough
	case false:
		fmt.Println("false still prints")
	}

	// #9
	favSport := "badminton"
	switch favSport {
	case "badminton":
		fmt.Println("CORRECT!")
	case "football":
		fmt.Println("wrong")
	default:
		fmt.Println("default!") // prints when no case is correct unless fallthrough
	}

	// #10 -> true, false, true, true, false
}