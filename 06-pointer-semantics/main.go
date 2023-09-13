package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi, my name is bikibiki slim", p.first, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Shady",
		age:   25,
	}

	saySomething(&p1)
	// does not work
	// saySomething(p1)
}
