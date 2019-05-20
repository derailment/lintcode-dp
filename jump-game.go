package main

/**
 * @param A: A list of integers
 * @return: A boolean
 */
func canJump(A []int) bool {
	n := len(A)
	f := make([]bool, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			f[i] = true
		} else {
			isReach := false
			for j := 0; j < i; j++ {
				isReach = isReach || f[j] && (A[j] >= i-j)
			}
			f[i] = isReach
		}
	}
	return f[n-1]
}
