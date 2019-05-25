package main

/**
 * @param A: An integer array
 * @param k: A positive integer (k <= length(A))
 * @param target: An integer
 * @return: An integer
 */
func kSum(A []int, k int, target int) int {
	n := len(A)
	f := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			f[i][j] = make([]int, target+1)
			for s := 0; s <= target; s++ {
				if s == 0 && j == 0 {
					f[i][j][s] = 1
				} else if s == 0 && j > 0 {
					f[i][j][s] = 0
				} else if i == 0 || j == 0 {
					f[i][j][s] = 0
				} else {
					if s >= A[i-1] {
						f[i][j][s] = f[i-1][j][s] + f[i-1][j-1][s-A[i-1]]
					} else {
						f[i][j][s] = f[i-1][j][s]
					}
				}
			}
		}
	}
	return f[n][k][target]
}
