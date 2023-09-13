package main

import (
	"fmt"
	"math/rand"

	"github.com/jsalinlee/puppy"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	var coolStr string = "I'm a pretty dope string."

	const reliableInt = 42
	thePerfectPair := "42"

	fmt.Printf("A cool string would be \"%v\"\n", coolStr)
	fmt.Printf("A great number is %v\n", reliableInt)
	fmt.Printf("The perfect pair is %v of type %T\n", thePerfectPair, thePerfectPair)

	// go mod practice
	puppy.From13()

	s1 := puppy.Bark()
	s2 := puppy.Barks()

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(puppy.BigBarks())

	// rand package
	x := rand.Intn(400)

	fmt.Printf("x = %v\n", x)

	if x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x <= 200 {
		fmt.Println("Between 101 and 200")
	} else if x <= 250 {
		fmt.Println("Between 201 and 250")
	} else {
		fmt.Println("More than 250")
	}

	switch {
	case x <= 100:
		fmt.Println("Between 0 and 100")
	case x <= 200:
		fmt.Println("Between 101 and 200")
	case x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("More than 250")
	}
}
