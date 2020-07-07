package main

import (
	"fmt"
	"sync"
	"runtime"
)

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

}