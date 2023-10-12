package dog

import (
	"fmt"
	"testing"
)

type test struct {
	data int
	want int
}

func TestYears(t *testing.T) {
	tests := []test{
		test{4, 28},
		test{8, 56},
		test{1, 7},
		test{5, 35},
	}

	for _, v := range tests {
		res := Years(v.data)
		if res != v.want {
			t.Error("want", v.want, "got", res)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	tests := []test{
		test{4, 28},
		test{8, 56},
		test{1, 7},
		test{5, 35},
	}

	for _, v := range tests {
		res := YearsTwo(v.data)
		if res != v.want {
			t.Error("want", v.want, "got", res)
		}
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(4))
	// Output:
	// 28
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(6))
	// Output:
	// 42
}
