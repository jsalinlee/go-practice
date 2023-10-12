package word

import (
	"fmt"
	"practicemodule/18-testing/exercise-2/quote"
	"testing"
)

type test struct {
	data   string
	target string
	want   int
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "want 3")
	}
}
func TestUseCount(t *testing.T) {
	tests := []test{
		test{"and and to and the of many the to too and", "and", 4},
		test{"no words here to search for", "for", 1},
		test{"no words here to search for unless you need to look for a word", "to", 2},
	}
	for _, v := range tests {
		m := UseCount(v.data)
		got := m[v.target]
		if got != v.want {
			t.Error("want", v.want, "got", got)
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
