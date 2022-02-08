package old

/*
204. 计数质数
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/
func countPrimes(n int) int {
	primes := make([]bool, n)
	j, count := 0, 0
	for i := 2; i < n; i++ {
		if !primes[i] {
			for j = 2; j*i < n; j++ {
				primes[j*i] = true
			}
			count++
		}
	}
	return count
}
