package main

/**
 * @param s: A string
 * @param p: A string includes "." and "*"
 * @return: A boolean
 */
func isMatch2(s string, p string) bool {
	n := len(s)
	m := len(p)
	f := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]bool, m+1)
		for j := 0; j < m+1; j++ {
			if i == 0 && j == 0 {
				f[i][j] = true
			} else if i == 0 && j < 3 {
				f[i][j] = p[j-1] == '*'
			} else if i == 0 && j >= 3 {
				f[i][j] = p[j-1] == '*' && f[i][j-2]
			} else if i > 0 && j == 0 {
				f[i][j] = false
			} else { // i > 0 && j > 0
				f[i][j] = (p[j-1] == '.' && f[i-1][j-1]) || (p[j-1] == s[i-1] && f[i-1][j-1]) || (p[j-1] == '*' && j >= 2 && ((p[j-2] == '.' || p[j-2] == s[i-1]) && f[i-1][j] || f[i][j-2]))
			}
		}
	}
	return f[n][m]
}
