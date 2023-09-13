package main

import "fmt"

func main() {

	// slice initialization with specified size
	xi := make([]int, 10)

	for i := 0; i < 10; i++ {
		xi[i] = i
	}

	// slice with specified size, but no initialization
	xi2 := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		xi2 = append(xi2, i)
	}

	for _, v := range xi {
		fmt.Printf("%v is of type %T\n", v, v)
	}

	// slice manipulation
	xi3 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi3[:5])
	fmt.Println(xi3[5:])
	fmt.Println(xi3[2:7])
	fmt.Println(xi3[1:6])

	// push to slice(append)
	xi4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi4)
	xi4 = append(xi4, 52)
	fmt.Println(xi4)
	xi4 = append(xi4, 53, 54, 55)
	fmt.Println(xi4)
	y := []int{56, 57, 58, 59, 60}
	xi4 = append(xi4, y...)
	fmt.Println(xi4)

	// combining slices
	xi5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi5)
	xi5 = append(xi5[:3], xi5[6:]...)
	fmt.Println(xi5)

	// slice of strings
	xs := make([]string, 0, 50)
	xs = append(xs, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	for i := 0; i < len(xs); i++ {
		fmt.Printf("%v - %v\n", i, xs[i])
	}
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	// 2D slice
	xxs := [][]string{jb, jm}
	for i1, xs := range xxs {
		fmt.Printf("%v:\t", i1)
		for i2, s := range xs {
			fmt.Printf("%v - %v\t", i2, s)
		}
		fmt.Println("")
	}
}
