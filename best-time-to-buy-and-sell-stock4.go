package main

/**
 * @param K: An integer
 * @param prices: An integer array
 * @return: Maximum profit
 */
func maxProfit(K int, prices []int) int {
	if K > len(prices)/2 {
		return maxProfit2(prices)
	}
	return maxProfit4(K, prices)
}

/**
 * @param K: An integer
 * @param prices: An integer array
 * @return: Maximum profit
 */
func maxProfit4(K int, prices []int) int {
	n := len(prices)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2*K+2)
		for j := 1; j <= 2*K+1; j++ {
			if i == 0 {
				f[i][j] = 0
			} else if i == 1 {
				f[i][j] = 0
			} else {
				profit := prices[i-1] - prices[i-2]
				if j == 1 {
					f[i][j] = f[i-1][j]
				} else if j%2 == 1 {
					f[i][j] = max(f[i-1][j], f[i-1][j-1]+profit)
				} else {
					f[i][j] = max(f[i-1][j]+profit, f[i-1][j-1])
				}
			}
		}
	}
	s := make([]int, K+1)
	for i := 0; i <= K; i++ {
		s[i] = f[n][2*K+1]
	}
	return maxArr(s)
}
