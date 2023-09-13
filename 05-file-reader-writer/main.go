package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	wordFreq := make(map[string]int)

	f, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	// Create a scanner to read the file
	s := bufio.NewScanner(f)
	s.Split((bufio.ScanWords))
	for s.Scan() {
		wordFreq[s.Text()]++
	}

	if err := s.Err(); err != nil {
		return
	}

	for k, v := range wordFreq {
		fmt.Printf("%v: %v\n", k, v)
	}
}
