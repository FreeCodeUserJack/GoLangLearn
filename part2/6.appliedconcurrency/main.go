package main

import (
	"fmt"
	"sync"
)

func main() {
	// factorial with pipeline pattern concurrently and in parallel
	res := []chan int{}
	for i := 2; i < 102; i++ {
		res = append(res, factorial(i))
	}
	for _, val := range res {
		fmt.Print(<-val, " ")
	}
	fmt.Println()

	// fan in / fan out example
	in := gen(2, 3)
	in2 := gen(5, 6)
	c1 := sq(in)
	c2 := sq(in2)
	fmt.Printf("C1=C2? - %t, memory addresses - %p, %p, values: %+v, %+v\n", c1 == c2, &c1, &c2, c1, c2)

	for n := range merge(c1, c2) {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
}

func factorial(n int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for n > 1 {
			total *= n
			n--
		}
		out <- total
		close(out)
	}()

	return out
}

// gen
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// intermediary processing
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// fan out
func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Printf("TYPE: %T\n", cs)
	wg.Add(len(cs))
	fmt.Println("number of channels to fan in: ", len(cs))
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				// fmt.Println(n)
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}