package main

import (
	"fmt"
	"practicemodule/18-testing/exercise-2/quote"
	"practicemodule/18-testing/exercise-2/word"
)

func main() {

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
	fmt.Println(word.Count(quote.SunAlso))
}
