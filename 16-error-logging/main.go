package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	n, err := os.Open("names.txt")
	if err != nil {
		log.Println("Error - open file", err)
		return
	}
	defer n.Close()
}
