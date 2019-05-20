package main

/**
 * @param prices: Given an integer array
 * @return: Maximum profit
 */
func maxProfit2(prices []int) int {
	maxProfit := 0
	for j := 0; j < len(prices)-1; j++ {
		profit := prices[j+1] - prices[j]
		if profit > 0 {
			maxProfit += profit
		}
	}
	return maxProfit
}
