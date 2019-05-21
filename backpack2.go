package main

/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @param V: Given n items with value V[i]
 * @return: The maximum value
 */
func backPackII(m int, A []int, V []int) int {
	n := len(A)
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
		for w := 0; w < m+1; w++ {
			if w == 0 {
				f[i][w] = 0
			} else if i == 0 {
				f[i][w] = -1
			} else if w >= A[i-1] && f[i-1][w-A[i-1]] != -1 {
				f[i][w] = max(f[i-1][w], f[i-1][w-A[i-1]]+V[i-1])
			} else {
				f[i][w] = f[i-1][w]
			}
		}
	}
	return maxArr(f[n])
}
