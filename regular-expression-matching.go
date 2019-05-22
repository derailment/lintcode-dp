package main

import (
	"fmt"
)

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

		}
	}
	return f[n][m]
}

func main() {
	fmt.Println(isMatch2("aa", "a"))
}
