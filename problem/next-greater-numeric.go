package problem

import (
	"sort"
	"sync"
)

// https://leetcode.com/problems/next-greater-numerically-balanced-number

// NextGreaterNumber finds the smallest numerically balanced number greater than n.
// A numerically balanced number is defined such that for every digit d in the number,
// the digit d appears exactly d times in the number.
//
// Time Complexity: O(k * log k) where k is the number of digits in the next balanced number.
func NextGreaterNumber(n int) int {
	for x := n + 1; ; x++ {
		if isNumericallyBalanced(x) {
			return x
		}	
	}
}

func isNumericallyBalanced(num int) bool {
	var count [10]int
	if num == 0 {
		count[0] = 1
	}
	for v := num; v > 0; v /= 10 {
		count[v%10]++
	}

	// digit 0 must not appear
	if count[0] > 0 {
		return false
	}
	// for digits 1-9, either does not appear
	// or appears exactly d times
	for d := 1; d <= 9; d++ {
		if count[d] != 0 && count[d] != d {
			return false
		}
	}
	return true
}

var (
    precomputedBalanced []int
    precomputeOnce      sync.Once
)

// Generation approach (precomputation) - generates all numerically balanced numbers
// up to a certain limit and stores them in a sorted slice for quick lookup.
// This approach is more efficient for multiple queries but requires more initial computation.
func NextGreaterNumberPrecomputed(n int) int {
	arr := generateNumericallyBalancedNumbers()
    // find first element strictly greater than n
    idx := sort.SearchInts(arr, n+1)
    if idx < len(arr) {
        return arr[idx]
    }
    return -1
}

func generateNumericallyBalancedNumbers() []int {
	precomputeOnce.Do(func() {
        limit := 1000000
        tmp := make([]int, 0, 100) // expected to be small
        for num := 1; num <= limit; num++ {
            if isNumericallyBalanced(num) {
                tmp = append(tmp, num)
            }
        }
        sort.Ints(tmp)
        precomputedBalanced = tmp
    })
    return precomputedBalanced
}
