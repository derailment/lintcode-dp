package main

/**
 * @param n: An integer
 * @return: A boolean which equals to true if the first player will win
 */
func firstWillWin(n int) bool {
	f := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		if i == 0 {
			f[i] = false
		} else if i == 1 {
			f[i] = true
		} else if i == 2 {
			f[i] = true
		} else {
			if f[i-1] == true && f[i-2] == true {
				f[i] = false
			}
			if f[i-1] == false && f[i-2] == true {
				f[i] = true
			}
			if f[i-1] == true && f[i-2] == false {
				f[i] = true
			}
			if f[i-1] == false && f[i-2] == false {
				f[i] = true
			}
		}
	}
	return f[n]
}
