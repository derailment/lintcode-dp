package main

import (
	"math"
	"strconv"
)

/**
 * @param s: a message being encoded
 * @return: an integer
 */
func numDecodings2(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		if i == 0 {
			f[i] = 1
		} else if i == 1 {
			f[i] = oneDigitWays(s[i-1 : i])
		} else {
			f[i] = 0
			f[i] += f[i-1] * oneDigitWays(s[i-1:i])
			f[i] += f[i-2] * twoDigitWays(s[i-2:i])
		}
	}
	return f[n] % int(math.Pow10(9)+7)
}

func oneDigitWays(s string) int {
	i, _ := strconv.Atoi(s)
	if s == "*" {
		return 9
	} else if 1 <= i && i <= 9 {
		return 1
	}
	return 0
}

func twoDigitWays(s string) int {
	i1, _ := strconv.Atoi(s[1:2])
	i0, _ := strconv.Atoi(s[0:1])
	if s == "**" {
		return 15
	} else if s[0] == '*' && 0 <= i1 && i1 <= 6 {
		return 2
	} else if s[0] == '*' && 7 <= i1 && i1 <= 9 {
		return 1
	} else if s[1] == '*' && i0 == 1 {
		return 9
	} else if s[1] == '*' && i0 == 2 {
		return 6
	} else if i0 == 1 && 0 <= i1 && i1 <= 9 {
		return 1
	} else if i0 == 2 && 0 <= i1 && i1 <= 6 {
		return 1
	}
	return 0
}
