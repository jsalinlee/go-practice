package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		want int
	}

	tests := []test{
		test{[]int{22, 20}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{-3, 5}, 2},
		test{[]int{-1, 0, 1}, 0},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.want {
			t.Error("Expected", v.want, "got", x)
		}
	}
}
