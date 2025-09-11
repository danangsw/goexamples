package testing_test

import (
	"reflect"
	"testing"

	problem "../problem"
)

func TestStringSplit(t *testing.T) {
	testCases := []struct {
		s      string
		del    rune
		expect []string
	}{
		{
			s:      "Hello World Let's Go!",
			del:    ' ',
			expect: []string{"Hello", "World", "Let's", "Go!"},
		},
		{
			s:      "apple,banana,orange",
			del:    ',',
			expect: []string{"apple", "banana", "orange"},
		},
		{
			s:      "one two three",
			del:    'x',
			expect: []string{"one two three"},
		},
		{
			s:      "",
			del:    ' ',
			expect: []string{},
		},
		{
			s:      "   leading spaces",
			del:    ' ',
			expect: []string{"", "", "", "leading", "spaces"},
		},
		{
			s:      "trailing spaces   ",
			del:    ' ',
			expect: []string{"trailing", "spaces", "", "", ""},
		},
		{
			s:      "  multiple   spaces  ",
			del:    ' ',
			expect: []string{"", "", "multiple", "", "", "spaces", "", ""},
		},
		{
			s:      "你好世界",
			del:    ' ',
			expect: []string{"你好世界"},
		},
		{
			s:      "你好 世界",
			del:    ' ',
			expect: []string{"你好", "世界"},
		},
		{
			s:      "a,b,c,d,e",
			del:    ',',
			expect: []string{"a", "b", "c", "d", "e"},
		},
		{
			s:      "abcde",
			del:    ' ',
			expect: []string{"abcde"},
		},
	}

	for _, tc := range testCases {
		result := problem.StringSplit(tc.s, tc.del)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("For s=%q and del=%q, expected %v, but got %v", tc.s, tc.del, tc.expect, result)
		}
	}
}
