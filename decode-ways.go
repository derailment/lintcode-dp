package main

import (
	"strconv"
)

/**
 * @param s: a string,  encoded message
 * @return: an integer, the number of ways decoding
 */
func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	f := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			if isDecode(s[0:1]) {
				f[i] = 1
			} else {
				f[i] = 0
			}
		} else if i == 1 {
			if isDecode(s[0:2]) {
				f[i] = 1
			} else {
				f[i] = 0
			}
			if isDecode(s[1:2]) {
				f[i] += f[i-1]
			}
		} else {
			f[i] = 0
			if isDecode(s[i : i+1]) {
				f[i] += f[i-1]
			}
			if isDecode(s[i-1 : i+1]) {
				f[i] += f[i-2]
			}
		}
	}
	return f[n-1]
}

func isDecode(s string) bool {
	if len(s) == 2 && s[0] == '0' {
		return false
	}
	d, _ := strconv.Atoi(s)
	if d == 0 {
		return false
	}
	if d > 26 {
		return false
	}
	return true
}
