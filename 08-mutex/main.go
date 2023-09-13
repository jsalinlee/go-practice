package main

import (
	"fmt"
	"sync"
)

func main() {

	gs := 100
	incrementer := 0

	var wg sync.WaitGroup
	wg.Add(gs)

	var mux sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mux.Lock()
			incrementer++
			fmt.Println(incrementer)
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)
}
