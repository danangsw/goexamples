package problem

// https://leetcode.com/problems/reverse-bits/

// pre computed table for reversed 8-bit value
var table [256]int

// using table[0..255] to store reversed bits of 0..255
// where table[i] = reverseBits(i)
func init() {
	for i := 0; i < 256; i++ {
		table[i] = reverse8Bits(i)
	}
}

// ReverseBits reverses the bits of a given 32 bits unsigned integer.
// Time Complexity: O(1) - constant time due to fixed number of operations
// Space Complexity: O(1) - constant space for the lookup table	
func ReverseBits(n int) int {
	b0 := int(n & 0xff)
    b1 := int((n >> 8) & 0xff)
    b2 := int((n >> 16) & 0xff)
    b3 := int((n >> 24) & 0xff)

    return (table[b0] << 24) | (table[b1] << 16) | (table[b2] << 8) | table[b3]
}

// reverse8Bits reverses the bits of an 8-bit integer.
func reverse8Bits(b int) int {
	var rev int = 0
	for i := 0; i < 8; i++ {
		rev = (rev << 1) | (b & 1)
		b >>= 1
	}
	return rev
}

// ReverseBitsIterative reverses the bits of a given 32 bits unsigned integer using iterative method.
// Time Complexity: O(1) - constant time due to fixed number of operations
// Space Complexity: O(1) - constant space
func ReverseBitsIterative(n int) int {
	var rev int = 0
	for i := 0; i < 32; i++ {
		rev = (rev << 1) | (n & 1)
		n >>= 1
	}
	return rev
}
