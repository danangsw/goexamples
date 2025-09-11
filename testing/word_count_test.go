package testing_test

import (
    "reflect"
    "testing"
	problem "../problem"
)

func TestWordCount(t *testing.T) {
    testCases := []struct {
        s      string
        expect map[string]int
    }{
        {
            s:      "Hello World Let's Go!",
            expect: map[string]int{"Hello": 1, "World": 1, "Let's": 1, "Go!": 1},
        },
        {
            s:      "apple banana orange",
            expect: map[string]int{"apple": 1, "banana": 1, "orange": 1},
        },
        {
            s:      "one two three one two one",
            expect: map[string]int{"one": 3, "two": 2, "three": 1},
        },
        {
            s:      "",
            expect: map[string]int{},
        },
        {
            s:      "   leading spaces",
            expect: map[string]int{"leading": 1, "spaces": 1},
        },
        {
            s:      "trailing spaces   ",
            expect: map[string]int{"trailing": 1, "spaces": 1},
        },
        {
            s:      "  multiple   spaces  ",
            expect: map[string]int{"multiple": 1, "spaces": 1},
        },
        {
            s:      "你好世界",
            expect: map[string]int{"你好世界": 1},
        },
        {
            s:      "你好 世界",
            expect: map[string]int{"你好": 1, "世界": 1},
        },
        {
            s:      "a b c d e",
            expect: map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 1},
        },
        {
            s:      "the quick brown fox jumps over the lazy dog",
            expect: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1},
        },
    }

    for _, tc := range testCases {
        result := problem.WordCount(tc.s)
        if !reflect.DeepEqual(result, tc.expect) {
            t.Errorf("For s=%q, expected %v, but got %v", tc.s, tc.expect, result)
        }
    }
}