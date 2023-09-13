package main

import "fmt"

func main() {

	a := []int{1, 2}
	// b := a
	// a[0] = 3
	b := make([]int, 2)
	copy(b, a)
	a[0] = 3
	fmt.Println(a)
	fmt.Println(b)
}
