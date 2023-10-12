package main

import (
	"fmt"
	"practicemodule/18-testing/acdc"
)

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 7 =", mySum(4, 7))
	fmt.Println("5 + 3 =", mySum(5, 3))
	fmt.Println(acdc.Sum(2, 34, 2, 6, 7, 2, 9, 5))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
