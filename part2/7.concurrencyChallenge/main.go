package main

import (
	"fmt"
	"time"
)

// not working, might come back later
func main() {

	in := gen()

	f := factorial(in)

	for n := range merge(f) {
		fmt.Println(n)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("func end")
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) []chan int {
	res := []chan int{}
	for n := range in {
		go func(i int) {
			out := make(chan int)
			out <- fact(i)
			close(out)
			res = append(res, out)
		}(n)
	}

	return res
}

func merge(chs []chan int) chan int {
	out := make(chan int)
	for _, ch := range chs {
		go func(c chan int) {
			out <- <-c
		}(ch)
	}
	close(out)
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/