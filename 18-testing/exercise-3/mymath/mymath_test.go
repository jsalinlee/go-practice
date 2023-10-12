package mymath

import (
	"fmt"
	"testing"
)

type test struct {
	data []int
	want float64
}

func TestCenteredAvg(t *testing.T) {
	tests := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}
	for _, v := range tests {
		c := CenteredAvg(v.data)
		if c != v.want {
			t.Error("got", c, "want", v.want)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{9000, 4, 10, 8, 6, 12}))
	// Output:
	// 9
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{123, 744, 140, 200})
	}
}
