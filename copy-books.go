package main

/**
 * @param pages: an array of integers
 * @param k: An integer
 * @return: an integer
 */
func copyBooks(pages []int, k int) int {
	n := len(pages)
	f := make([][]int, n+1)
	costSum := make([]int, n)
	costSum[0] = pages[0]
	for j := 0; j < n+1; j++ {
		if j != 0 && j != 1 {
			costSum[j-1] = costSum[j-2] + pages[j-1]
		}
		f[j] = make([]int, k+1)
		for p := 1; p < k+1; p++ {
			if j == 0 {
				f[j][p] = 0
			} else if p == 1 {
				f[j][p] = costSum[j-1]
			} else {
				minCost := int(^uint(0) >> 1)
				for i := 1; i <= j; i++ {
					partitionCost := max(f[i][p-1], costSum[j-1]-costSum[i-1])
					if partitionCost < minCost {
						minCost = partitionCost
					}
				}
				f[j][p] = minCost
			}
		}
	}
	return f[n][k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
