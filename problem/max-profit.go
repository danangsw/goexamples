package problem

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

// return the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.

func MaxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
