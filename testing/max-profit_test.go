package testing_test

import (
	"testing"

	problem "../problem"
)

func TestMaxProfit(t *testing.T) {
    tests := []struct {
        name     string
        prices   []int
        expected int
    }{
        {
            name:     "Typical case",
            prices:   []int{7, 1, 5, 3, 6, 4},
            expected: 5, // Buy at 1, sell at 6
        },
        {
            name:     "Prices always decreasing",
            prices:   []int{7, 6, 4, 3, 1},
            expected: 0, // No profit possible
        },
        {
            name:     "Prices always increasing",
            prices:   []int{1, 2, 3, 4, 5},
            expected: 4, // Buy at 1, sell at 5
        },
        {
            name:     "Single day",
            prices:   []int{5},
            expected: 0, // Not enough data to sell
        },
        {
            name:     "Empty array",
            prices:   []int{},
            expected: 0, // No data
        },
        {
            name:     "Two days profit",
            prices:   []int{1, 5},
            expected: 4,
        },
        {
            name:     "Two days loss",
            prices:   []int{5, 1},
            expected: 0,
        },
        {
            name:     "Multiple peaks and valleys",
            prices:   []int{3, 2, 6, 1, 3},
            expected: 4, // Buy at 2, sell at 6
        },
        {
            name:     "Profit at the end",
            prices:   []int{5, 4, 3, 2, 1, 10},
            expected: 9, // Buy at 1, sell at 10
        },
        {
            name:     "Large input",
            prices:   generateLargeInput(100000),
            expected: 99999, // Buy at 1, sell at 100000
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := problem.MaxProfit(tt.prices)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}

// Helper to generate large increasing input
func generateLargeInput(n int) []int {
    prices := make([]int, n)
    for i := 0; i < n; i++ {
        prices[i] = i + 1
    }
    return prices
}