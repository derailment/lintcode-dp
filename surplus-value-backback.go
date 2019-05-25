package main

import (
	"sort"
)

/**
 * @param k1: The coefficient of A
 * @param k2: The  coefficient of B
 * @param c: The volume of backpack
 * @param n: The amount of A
 * @param m: The amount of B
 * @param a: The volume of A
 * @param b: The volume of B
 * @return: Return the max value you can get
 */
func getMaxValue(k1 int, k2 int, c int, n int, m int, a []int, b []int) int64 {
	sort.Ints(a)
	sort.Ints(b)
	f := make([][]int, n+1)
	loadA := 0
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
		if i != 0 {
			loadA += a[i-1]
		}
		loadB := 0
		for j := 0; j <= m; j++ {
			if j != 0 {
				loadB += b[j-1]
			}
			remaining := c - loadA - loadB
			if i == 0 && j == 0 {
				f[i][j] = 0
			} else if i == 0 && remaining >= 0 {
				f[i][j] = f[i][j-1] + k2*remaining
			} else if j == 0 && remaining >= 0 {
				f[i][j] = f[i-1][j] + k1*remaining
			} else if remaining >= 0 {
				f[i][j] = max(f[i-1][j]+k1*remaining, f[i][j-1]+k2*remaining)
			} else if i == 0 && j > 0 && remaining < 0 {
				f[i][j] = f[i][j-1]
			} else if i > 0 && j == 0 && remaining < 0 {
				f[i][j] = f[i-1][j]
			} else if remaining < 0 {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return int64(max2dArr(f))
}

func max2dArr(arr [][]int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] > max {
				max = arr[i][j]
			}
		}
	}
	return max
}
