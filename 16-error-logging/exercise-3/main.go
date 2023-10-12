package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("We're going down down - %v", c.info)
}

func main() {
	ce := customErr{"in an earlier round"}
	foo(ce)
}

func foo(e error) {
	log.Println("Here's the error - ", e, "\n", e.(customErr).info)
}
