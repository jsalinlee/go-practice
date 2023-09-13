package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Goroutine numero uno")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine numero segundo uno")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Goroutines should be running, if not you forgot the waitgroups")
	fmt.Println(runtime.NumCPU())
}
