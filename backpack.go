package main

import (
	"fmt"
)

/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @return: The maximum size
 */
func backPack(m int, A []int) int {
	n := len(A)
	f := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]bool, m+1)
		for j := 0; j < m+1; j++ {
			if j == 0 {
				f[i][j] = true
			} else if i == 0 {
				f[i][j] = false
			} else {
				if j-A[i-1] >= 0 {
					f[i][j] = f[i-1][j] || f[i-1][j-A[i-1]]
				} else {
					f[i][j] = f[i-1][j] || false
				}
			}
		}
	}
	return maxWeight(f)
}

func maxWeight(f [][]bool) int {
	for j := len(f[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(f); i++ {
			if f[i][j] == true {
				return j
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(backPack(1, []int{3, 4, 8, 5}))
}
