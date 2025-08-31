package testing_test

import (
	"fmt"
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
		t.Run(fmt.Sprintf("IsPalindrome(%d)", c.in), func(t *testing.T) {
			result := problem.IsPalindrome(c.in)
			if result != c.want {
				t.Errorf("expected %v, got %v", c.want, result)
			}
		})
	}
}
