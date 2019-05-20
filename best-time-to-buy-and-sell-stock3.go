package main

import (
	"fmt"
)

/**
 * @param prices: Given an integer array
 * @return: Maximum profit
 */
func maxProfit3(prices []int) int {
	n := len(prices)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 6)
		for j := 1; j <= 5; j++ {
			if i == 0 {
				f[i][j] = 0
			} else if i == 1 {
				f[i][j] = 0
			} else {
				profit := prices[i-1] - prices[i-2]
				if j == 1 {
					f[i][j] = f[i-1][j]
				}
				if j == 3 || j == 5 {
					f[i][j] = maxNum(f[i-1][j], f[i-1][j-1]+profit)
				}
				if j == 2 || j == 4 {
					f[i][j] = maxNum(f[i-1][j]+profit, f[i-1][j-1])
				}
			}
		}
	}
	return maxArr([]int{f[n][1], f[n][3], f[n][5]})
}

func main() {
	fmt.Println(maxProfit3([]int{4, 4, 6, 1, 1, 4, 2, 5}))
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArr(s []int) int {
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
