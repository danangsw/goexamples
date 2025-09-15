package testing_test

import (
	"testing"
	problem "../problem"
)

func TestLengthOfLastWord1(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{s: "Hello World", expect: 5},
		{s: "   fly me   to   the moon  ", expect: 4},
		{s: "luffy is still joyboy", expect: 6},
		{s: "word", expect: 4},
		{s: " word", expect: 4},
		{s: "word ", expect: 4},
		{s: " word ", expect: 4},
		{s: "a", expect: 1},
		{s: " a", expect: 1},
		{s: "a ", expect: 1},
		{s: " a ", expect: 1},
		{s: "  a  ", expect: 1},
		{s: " b   a    ", expect: 1},
	}

	for _, tc := range testCases {
		result := problem.LengthOfLastWord1(tc.s)
		if result != tc.expect {
			t.Errorf("LengthOfLastWord1: For s=%q, expected %v, but got %v", tc.s, tc.expect, result)
		}
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{s: "Hello World", expect: 5},
		{s: "   fly me   to   the moon  ", expect: 4},
		{s: "luffy is still joyboy", expect: 6},
		{s: "word", expect: 4},
		{s: " word", expect: 4},
		{s: "word ", expect: 4},
		{s: " word ", expect: 4},
		{s: "a", expect: 1},
		{s: " a", expect: 1},
		{s: "a ", expect: 1},
		{s: " a ", expect: 1},
		{s: "  a  ", expect: 1},
		{s: " b   a    ", expect: 1},
	}

	for _, tc := range testCases {
		result := problem.LengthOfLastWord2(tc.s)
		if result != tc.expect {
			t.Errorf("LengthOfLastWord2: For s=%q, expected %v, but got %v", tc.s, tc.expect, result)
		}
	}
}
