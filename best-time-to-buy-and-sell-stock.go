package main

/**
 * @param prices: Given an integer array
 * @return: Maximum profit
 */
func maxProfit1(prices []int) int {
	maxProfit := 0
	for j := 1; j < len(prices); j++ {
		for i := 0; i < j; i++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}

		}
	}
	return maxProfit
}
