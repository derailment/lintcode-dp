package lintcode

/**
 * @param coins: a list of integer
 * @param amount: a total amount of money amount
 * @return: the fewest number of coins that you need to make up
 */
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	const MAXINT = int(^uint(0) >> 1)
	f := make([]int, amount+1)
	for i := range f {
		// f(x == 0) = 0
		if i == 0 {
			f[i] = 0
		}
		// f(x > 0) = min {f(x-a)+1, f(x-b)+1,... ,f(x-c)+1}
		if i > 0 {
			coinNum := MAXINT
			for _, c := range coins {
				if c == 0 {
					continue
				}
				if i-c < 0 {
					continue
				}
				if f[i-c] == MAXINT {
					continue
				}
				tryCoinNum := f[i-c] + 1
				if tryCoinNum < coinNum {
					coinNum = tryCoinNum
				}
			}
			f[i] = coinNum
		}
	}
	if f[amount] == MAXINT {
		return -1
	}
	return f[amount]
}
