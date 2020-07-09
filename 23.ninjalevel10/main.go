package main

import (
	// "time"
	"runtime"
	"fmt"
	// "sync"
)

func main() {
	// #1 make code work
	c := make(chan int)
	defer close(c)
	go func(){c <- 42}()
	fmt.Println(<-c)

	// #2
	cs := make(chan int)
	defer close(cs)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("-------------\n")
	fmt.Printf("cs\t%T\n", cs)

	// #3
	cg := gen()
	receive(cg) // has to be goroutine else gen() will block indefinitely until there is a receiver
	fmt.Println("#3 done")

	// #4
	q := make(chan int)
	cc := gen2(q)

	receive2(cc, q)
	fmt.Println("#4 done")

	// #5
	co := make(chan int)
	go func(){
		co <- 25
	}()
	v, ok := <-co
	fmt.Println(v, ok)
	close(co)
	v, ok = <-co
	fmt.Println(v, ok)
	fmt.Println("#5 done")

	// #6
	c1 := make(chan int)
	go func(c chan<- int){
		for i:=0; i<50; i++ {
			c <- i
		}
		close(c)
	}(c1)
	for v := range c1 {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n#6 done")

	// #7
	x := 10
	y := 10
	cbuf := gen7(x, y)

	for ki:=0; ki<x*y; ki++ { // not sure how to close() in gen7 so we don't use range
		fmt.Println(ki, <-cbuf)
	}
	// for v := range cbuf {
	// 	fmt.Println(v)
	// }
	fmt.Println("Routines: ", runtime.NumGoroutine())
}

// helper functions
// #3
func gen() <-chan int {
	c := make(chan int)
	go func(){
		for i:=0; i<100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for val := range c {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

// #4
func gen2(q chan<- int) <-chan int {
	c := make(chan int)
	go func(){
		for i:=0; i<55; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive2(c <-chan int, q <-chan int) {
	for {
		select {
		case val := <-c:
			fmt.Printf("%d ", val)
		case i, ok := <-q:
			fmt.Println("\nreceived quit signal", i, ok)
			return
		}
	}
}

// #7
func gen7(x, y int) <-chan int {
	// var wg sync.WaitGroup
	c := make(chan int)
	// wg.Add(10)
	for i:=0; i<x; i++ {
		go func() {
			for j:=0; j<y; j++ {
				c <- j
			}
			// wg.Done()
		}()
		fmt.Println("Routines (inside): ", runtime.NumGoroutine())
	}
	// wg.Wait()
	// close(c)
	return c
}