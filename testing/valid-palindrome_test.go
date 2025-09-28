package testing_test

import (
	"testing"

	problem "../problem"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// ✅ Basic cases
		{"Simple palindrome", "madam", true},
		{"Simple non-palindrome", "hello", false},

		// ✅ Case insensitivity
		{"Mixed case palindrome", "RaceCar", true},
		{"Mixed case non-palindrome", "RaceFast", false},

		// ✅ With punctuation and spaces
		{"Phrase with punctuation", "A man, a plan, a canal: Panama", true},
		{"Phrase with punctuation (non-palindrome)", "race a car", false},

		// ✅ Empty and single character
		{"Empty string", "", true},
		{"Single space", " ", true},
		{"Single letter", "a", true},
		{"Single digit", "7", true},

		// ✅ Numbers and alphanumerics
		{"Numeric palindrome", "12321", true},
		{"Alphanumeric palindrome", "1a2b2a1", true},
		{"Alphanumeric non-palindrome", "1a2b3a1", false},

		// ✅ Unicode characters
		{"Unicode palindrome", "😊a😊", true},                     // after cleaning: "a"
		{"Unicode non-palindrome", "😊a😢", true},                 // after cleaning: "a"
		{"Accented letters", "Ésope reste et se repose", false}, // after cleaning: "soperesteetserepose"

		// ✅ Symbols and mixed input
		{"Symbols only", "!@#$%^&*", true},     // after cleaning: ""
		{"Symbols with letters", "!a!", true},  // after cleaning: "a"
		{"Symbols with mismatch", "!a@", true}, // after cleaning: "a"

		// ✅ Performance edge
		{"Long palindrome", generatePalindrome(100000), true},
		{"Long non-palindrome", generateNonPalindrome(100000), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.ValidPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("Input: %q → got %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Helper to generate a long palindrome string
func generatePalindrome(n int) string {
	half := make([]rune, n/2)
	for i := range half {
		half[i] = 'a'
	}
	full := string(half) + string(half)
	return full
}

// Helper to generate a long non-palindrome string
func generateNonPalindrome(n int) string {
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = 'a'
	}
	runes[n-1] = 'b' // break symmetry
	return string(runes)
}
