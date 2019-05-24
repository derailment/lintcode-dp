package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch2("aa", "aa"))
}

/**
 * @param s: A string
 * @param p: A string includes "?" and "*"
 * @return: is Match?
 */

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	f := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]bool, m+1)
		for j := 0; j < m+1; j++ {
			if i == 0 && j == 0 {
				f[i][j] = true
			} else if i == 0 && j > 0 {
				f[i][j] = p[j-1] == '*' && f[i][j-1]
			} else if i > 0 && j == 0 {
				f[i][j] = false
			} else if i > 0 && j > 0 {
				f[i][j] = (p[j-1] == '?' && f[i-1][j-1]) || (p[j-1] == '*' && (f[i][j-1] || f[i-1][j])) || (p[j-1] == s[i-1] && f[i-1][j-1])
			}
		}
	}
	return f[n][m]
}
