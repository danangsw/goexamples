// https://leetcode.com/problems/sqrtx/
package problem

func Sqrtx(x int) int {
	if x <= 1 {
		return x
	}

	// Better initial guess - use floating point division
	result := float64(x) / 2.0
	tolerance := 1e-10  // Precision threshold for convergence
	maxIterations := 20 // Limit iterations to prevent infinite loops

	// Iterate until convergence or max iterations
	for i := 0; i < maxIterations; i++ {
		next := (result + float64(x)/result) / 2.0

		// Check convergence using tolerance instead of exact equality
		if abs(next-result) < tolerance {
			break
		}
		result = next
	}

	return int(result)
}

// Helper function for absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
