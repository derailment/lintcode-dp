package main

/**
 * @param costs: n x 3 cost matrix
 * @return: An integer, the minimum cost to paint all houses
 */
func minCost(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, 3)
		if i == 0 {
			f[i][0] = costs[i][0]
			f[i][1] = costs[i][1]
			f[i][2] = costs[i][2]
		} else {
			f[i][0] = minArr([]int{f[i-1][1], f[i-1][2]}) + costs[i][0]
			f[i][1] = minArr([]int{f[i-1][0], f[i-1][2]}) + costs[i][1]
			f[i][2] = minArr([]int{f[i-1][0], f[i-1][1]}) + costs[i][2]
		}
	}
	return minArr([]int{f[n-1][0], f[n-1][1], f[n-1][2]})
}

func minArr(arr []int) int {
	min := int(^uint(0) >> 1)
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
