package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

type key string

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var userID key = "userId"
	var fname key = "fname"
	ctx = context.WithValue(ctx, userID, 777)
	ctx = context.WithValue(ctx, fname, "James")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	var ch chan int = make(chan int)

	go func() {
		// simulate long task
		uid := ctx.Value("userID").(int)
		time.Sleep(4 * time.Second)

		// check not running in vain
		// if ctx.Done()
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()
	
	select {
	case <- ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}