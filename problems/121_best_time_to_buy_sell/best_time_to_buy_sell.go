// Package best_time_to_buy_sell
// difficulty: easy
// link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package best_time_to_buy_sell

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	buyPrice := prices[0]
	for _, price := range prices[1:] {
		if price <= buyPrice {
			buyPrice = price
			continue
		}
		profit := price - buyPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
