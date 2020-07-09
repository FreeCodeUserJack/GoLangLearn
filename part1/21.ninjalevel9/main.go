package main

import (
	"fmt"
	"sync"
	"runtime"
	// "time"
	"sync/atomic"
)

// #2
type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println(p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// #1
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Hello there")
	go func(){
		fmt.Println("Greetings")
		wg.Done()
	}()
	go func() {
		fmt.Println("Salutations")
		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	// #2
	p1 := person{
		first: "James",
	}
	saySomething(&p1) // cannot pass in p1, has to be pointer p1
	p1.speak()

	// #3, #4, #5
	var wg2 sync.WaitGroup
	var count int64
	// count := 0
	// var mu sync.Mutex
	wg2.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			// mu.Lock()
			// v := count
			// fmt.Println(count)
			// runtime.Gosched()
			// time.Sleep(time.Second)
			// v++
			// count = v
			// mu.Unlock()
			atomic.AddInt64(&count, 1)
			wg2.Done()
		}()
		fmt.Println(runtime.NumGoroutine())
	}
	wg2.Wait()
	fmt.Println(count)

	// #6
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// #7
	// teach a topic
}