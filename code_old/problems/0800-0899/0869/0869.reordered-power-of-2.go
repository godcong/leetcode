package _0869

var powerOf2Digits = initPowerOf2Digits()

func initPowerOf2Digits() map[[10]uint8]bool {
	ret := make(map[[10]uint8]bool)
	for n := 1; n <= 1e9; n <<= 1 {
		ret[countDigits(n)] = true
	}
	return ret
}

func countDigits(n int) (cnt [10]uint8) {
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}
	return
}

func reorderedPowerOf2(n int) bool {
	return powerOf2Digits[countDigits(n)]
}
