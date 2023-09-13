package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	bar(c)

	// go func(){
	// 	c <- 42
	// }()
	// fmt.Println(c)
	// fmt.Printf("%T\n", c)
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
