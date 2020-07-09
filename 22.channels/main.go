package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	go func(){
		c <- 42 // blocking call until 42 is received from c
	}()
	fmt.Println(<-c)

	bc := make(chan int, 2) // buffered channel
	bc <- 44
	bc <- 45
	fmt.Println(<-bc)
	close(bc)

	fmt.Printf("%T\t%T\n", c, bc)

	// directional chan
	dc := make(chan <- int, 2)
	dc <- 5
	// <- dc will throw error
	fmt.Printf("%T\n", dc)

	// using channels
	var wg sync.WaitGroup
	wg.Add(2)
	cc := make(chan int)
	go foo(cc, &wg)
	go bar(cc, &wg)
	wg.Wait()
	fmt.Println("Done")

	go foo2(cc)
	bar2(cc)
	fmt.Println("Done without waitgroup")

	// range over channel
	ch := make(chan int)
	go loop(ch)
	for val := range ch {
		fmt.Printf("%d ", val)
	}

	// select
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)
	defer close(e)
	defer close(o)
	defer close(q)

	go send(e, o, q)
	receive(e, o, q)
	fmt.Println("done select")

	// fanin
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send2(even, odd)
	go receive2(even, odd, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("fanin done")

	// fan out
	chan1 := make(chan int)
	chan2 := make(chan int)

	go populate(chan1)
	// go fanOutIn(chan1, chan2)
	go fanOutInThrottled(chan1, chan2)
	for v := range chan2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\nfanout/in done")
}


// helper functions
func foo(cc chan<- int, wg *sync.WaitGroup) {
	cc <- 5
	cc <- 6
	wg.Done()
}

func bar(cc <-chan int, wg *sync.WaitGroup) {
	fmt.Println(<-cc)
	fmt.Println(<-cc)
	wg.Done()
}

func foo2(cc chan<- int) {
	cc <- 55
	cc <- 66
}

func bar2(cc <-chan int) {
	fmt.Println(<-cc)
	fmt.Println(<-cc)
}

func loop(c chan<- int) {
	for i:=0; i<100; i++ {
		c <- i
	}
	close(c)
}

// select
func send(e, o, q chan<- int) {
	for i:=0; i < 100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v:= <-e:
			fmt.Println("even: ", v)
		case v := <-o:
			fmt.Println("odd: ", v)
		case _ = <-q:
			fmt.Println("q channel 0 means quit")
			return
		}
	}
}

// fan in
func send2(even, odd chan<- int) {
	for i:=0; i<10; i++ {
		if i % 2 == 0 {
			even <- i
		}else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive2(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()
	
	go func(){
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

// fanout
func populate(c chan int) {
	for i:=0; i<100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for i := range c1 {
		wg.Add(1)
		go func(v2 int){
			c2 <- timeConsumingTask(v2)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c2)
}

// limit goroutines to 10 at most
func fanOutInThrottled(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)
	for i:=0; i<goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int){
					c2 <- timeConsumingTask(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingTask(i int) int {
	time.Sleep(time.Millisecond)
	return i
}