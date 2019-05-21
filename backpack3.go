package main

import (
	"fmt"
)

/**
 * @param A: an integer array
 * @param V: an integer array
 * @param m: An integer
 * @return: an array
 */
func backPackIII(A []int, V []int, m int) int {
	n := len(A)
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
		for w := 0; w < m+1; w++ {
			if w == 0 {
				f[i][w] = 0
			} else if i == 0 {
				f[i][w] = -1
			} else if w >= A[i-1] && f[i][w-A[i-1]] != -1 {
				f[i][w] = max(f[i-1][w], f[i][w-A[i-1]]+V[i-1])
			} else {
				f[i][w] = f[i-1][w]
			}
		}
	}
	return maxArr(f[n])
}

func main() {
	fmt.Println(backPackIII([]int{2, 3, 5, 7}, []int{1, 5, 2, 4}, 10))
}
