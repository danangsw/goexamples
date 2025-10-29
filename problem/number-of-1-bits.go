package problem

// https://leetcode.com/problems/number-of-1-bits/

// HammingWeight returns the number of '1' bits in the binary representation of a given unsigned integer.
// Time Complexity: O(1) - constant time due to fixed number of operations
// Space Complexity: O(1) - constant space

func HammingWeight(num uint32) int {
    count := 0
    for num > 0 {
        count += int(num & 1)
        num >>= 1
    }
    return count
}