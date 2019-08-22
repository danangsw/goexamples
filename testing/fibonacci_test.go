package testing_test

import (
	"testing"
	samples "github.com/danangsw/goexamples/samples"
)

func TestFibonacci(t *testing.T)  {
	cases := []struct {
		in, want int
	} {
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for _, c := range cases {
		got := samples.Fibonacci(c.in)
		if got != c.want {
			t.Errorf("n (%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

