package main

import (
	"fmt"
)

/**
 * @param values: a vector of integers
 * @return: a boolean which equals to true if the first player will win
 */
func firstWillWin3(values []int) bool {
	n := len(values)
	f := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := i; j < n; j++ {
			if i == j {
				f[i][j] = values[i]
			} else {
				f[i][j] = max(values[i]-f[i+1][j], values[j]-f[i][j-1])
			}
		}
	}
	return f[0][n-1] >= 0
}

func main() {
	fmt.Println(firstWillWin3([]int{1, 20, 4}))
}
