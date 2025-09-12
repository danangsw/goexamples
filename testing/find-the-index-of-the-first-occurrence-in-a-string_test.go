package testing_test

import (
	"../problem"
	"testing"
)

func TestFindString(t *testing.T) {
    testCases := []struct {
        haystack string
        needle   string
        expect   int
    }{
        {
            haystack: "sadbutsad",
            needle:   "sad",
            expect:   0,
        },
        {
            haystack: "leetcode",
            needle:   "leeto",
            expect:   -1,
        },
        {
            haystack: "hello",
            needle:   "ll",
            expect:   2,
        },
        {
            haystack: "aaaaa",
            needle:   "bba",
            expect:   -1,
        },
        {
            haystack: "",
            needle:   "",
            expect:   0,
        },
        {
            haystack: "a",
            needle:   "",
            expect:   0,
        },
        {
            haystack: "",
            needle:   "a",
            expect:   -1,
        },
        {
            haystack: "abc",
            needle:   "c",
            expect:   2,
        },
        {
            haystack: "mississippi",
            needle:   "issip",
            expect:   4,
        },
        {
            haystack: "mississippi",
            needle:   "pi",
            expect:   9,
        },
        {
            haystack: "a",
            needle:   "a",
            expect:   0,
        },
        {
            haystack: "abc",
            needle:   "abcd",
            expect:   -1,
        },
        {
            haystack: "abcd",
            needle:   "abcde",
            expect:   -1,
        },
        {
            haystack: "abcde",
            needle:   "bc",
            expect:   1,
        },
        {
            haystack: "abcde",
            needle:   "cde",
            expect:   2,
        },
        {
            haystack: "abcde",
            needle:   "ab",
            expect:   0,
        },
        {
            haystack: "abcde",
            needle:   "e",
            expect:   4,
        },
        {
            haystack: "abcde",
            needle:   "f",
            expect:   -1,
        },
        {
            haystack: "abcdeabcde",
            needle:   "abcde",
            expect:   0,
        },
        {
            haystack: "abcdeabcde",
            needle:   "cdea",
            expect:   2,
        },
    }

    for _, tc := range testCases {
        result := problem.FindString(tc.haystack, tc.needle)
        if result != tc.expect {
            t.Errorf("For haystack=%q and needle=%q, expected %v, but got %v", tc.haystack, tc.needle, tc.expect, result)
        }
    }
}