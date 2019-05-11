package main

/**
 * @param obstacleGrid: A list of lists of integers
 * @return: An integer
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else if i == 0 && j == 0 {
				f[i][j] = 1
			} else if i == 0 {
				f[i][j] = f[i][j-1]
			} else if j == 0 {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = f[i][j-1] + f[i-1][j]
			}
		}
	}
	return f[m-1][n-1]
}
