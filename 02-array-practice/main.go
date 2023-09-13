package main

import "fmt"

func main() {

	// Array initialization
	ai := [5]int{}

	for i := 0; i < 5; i++ {
		ai[i] = i
	}
	// range over slice
	for _, v := range ai {
		fmt.Printf("%v is of type %T\n", v, v)
	}
}
