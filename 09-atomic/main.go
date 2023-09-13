package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	gs := 100
	var incrementer int64

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)
}
