package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// switch syntax
	for i := 0; i < 100; i++ {
		x, y := rand.Intn(10), rand.Intn(10)

		fmt.Printf("x: %v\t y: %v\n", x, y)
		if x < 4 && y < 4 {
			fmt.Println("x and y are both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y are both greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		} else if x != 5 {
			fmt.Println("y is not 5")
		}

		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		case x != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("no cases matched yo")
		}
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("x is 0")
		case 1:
			fmt.Println("x is 1")
		case 2:
			fmt.Println("x is 2")
		case 3:
			fmt.Println("x is 3")
		default:
			fmt.Println("x does not exist or is 4")
		}
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%#U\n", i)
	}
}
