package main

import (
	"sync/atomic"
	// "time"
	"sync"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	// go fmt.Println("here") // does not throw error, but you can't .Done() on it b/c it's built in func
	bar()

	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	wg.Wait()

	// race condition
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	// counter := 0
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	// var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		// go func(){
		// 	mu.Lock()
		// 	v := counter
		// 	// time.Sleep(time.Second)
		// 	runtime.Gosched()
		// 	v++
		// 	counter = v
		// 	mu.Unlock()
		// 	wg.Done()
		// }()
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}

func foo() {
	defer wg.Done()

	i := 0
	for i < 10 {
		fmt.Println("foo: ", i)
		i++
	}
}

func bar() {
	i := 0
	for i < 10 {
		fmt.Println("bar: ", i)
		i++
	}
}