package testing_test

import (
	"testing"
	samples "../samples"
)

func TestNthPrime(t *testing.T)  {
	cases := []struct {
		in, want int
	} {
		{0, 0},
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{6, 13},
		{7, 17},
		{8, 19},
		{9, 23},
		{10, 29},
		{25, 97},
		{40, 173},
	}
	for _, c := range cases {
		got := samples.NthPrime(c.in, false)
		if got != c.want {
			t.Errorf("n (%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

