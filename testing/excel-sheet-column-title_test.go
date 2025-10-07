package testing_test

import (
	"testing"
	problem "../problem"
)

// https://leetcode.com/problems/excel-sheet-column-title/description/

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		input int
		expected string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
		{702, "ZZ"},
		{701, "ZY"},
		{1402, "BAX"},
		{2804, "DCV"}
	}

	for _, test := range tests {
		result := problem.ConvertToTitle(test.input)
		if result != test.expected {
			t.Errorf("ConvertToTitle(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}
