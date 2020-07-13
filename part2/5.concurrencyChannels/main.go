package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// waitgroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i := 1
		for i <= 5 {
			fmt.Print(i, " ")
			i++
		}
		fmt.Println()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("waitgorup done")

	// mutex and atomic
	var res int
	var res2 int64
	var mu sync.Mutex
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go inc(&res, &res2, &wg2, &mu)
	go inc(&res, &res2, &wg2, &mu)
	wg2.Wait()
	fmt.Println(res, res2)

	// channels
	ch := make(chan int)
	go func() {
		ch <- 5
	}()
	fmt.Println(<- ch)

	// range over channels
	c1 := make(chan int)
	go func() {
		for j := 0; j < 50; j++ {
			c1 <- j
		}
		close(c1)
	}()
	for val := range c1 {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// mutex
func inc(x *int, x2 *int64, wg2 *sync.WaitGroup, mu *sync.Mutex) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		*x++
		fmt.Print(*x, " ")
		mu.Unlock()
		atomic.AddInt64(x2, 1)
	}
	fmt.Println()
	wg2.Done()
}

// channels
// func gen() <-chan int {
// 	res := make(chan int)
// 	for i:=0; i<50; i++ {
// 		res <- i
// 	}
// 	close(res)
// 	return res
// }