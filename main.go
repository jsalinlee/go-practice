package main

import (
	"fmt"

	"github.com/jsalinlee/puppy"
)

func main() {
	puppy.From13()

	s1 := puppy.Bark()
	s2 := puppy.Barks()

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
