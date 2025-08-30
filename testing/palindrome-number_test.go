package testing_test

import (
	"testing"

	problem "../problem"
)

func TestPalindromeNumber(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
		{0, true},
		{9, true},
	}

	for _, c := range cases {
		got := problem.IsPalindrome(c.in)
		if got != c.want {
			t.Errorf("IsPalindrome(%d) = %v; want %v", c.in, got, c.want)
		}
	}
}
