package main

import (
	"fmt"
	"sync"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go send(e, o)
	//receive

	go func(){
		for v := range e {
			fmt.Println(v)
		}
		wg.Done()
	}()

	go func(){
		for v := range o {
			fmt.Println(v)
		}
		wg.Done()
	}()
	
	wg.Wait()
	// for v := range e {
	// 	fmt.Println(v)
	// }
	// for p := range o {
	// 	fmt.Println(p)
	// }
}

//send
func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}


// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var ws sync.WaitGroup
// 	var mu sync.Mutex
// 	ws.Add(1)
// 	c := make(chan int)
// 	var v int64 = 0
// 	for i := 1; i < 11; i++ {
// 		go func(vv *int64) {
// 			for j := 1; j < 11; j++ {
// 				c <- j
// 			}
// 			mu.Lock()
// 			v++
// 			mu.Unlock()
// 		}(&v)
// 	}

// 	go func() {
// 		for {
// 			mu.Lock()
// 			if v == 10 {
// 				ws.Done()
// 				mu.Unlock()
// 				return
// 			}
// 			mu.Unlock()
// 			b := <-c
// 			fmt.Println(b)
// 		}
// 	}()
// 	ws.Wait()
// }
