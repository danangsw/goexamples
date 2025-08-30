package testing_test

import (
	"testing"
	samples "../samples"
)

func TestReverse(t *testing.T)  {
	cases := []struct {
		in, want string
	} {
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := samples.Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}