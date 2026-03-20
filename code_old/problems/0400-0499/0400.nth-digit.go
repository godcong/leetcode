package _400

import (
	"math"
	"sort"
)

func totalDigits(n int) int {
	digits := 0
	for curLength, curCount := 1, 9; curLength <= n; curLength++ {
		digits += curLength * curCount
		curCount *= 10
	}
	return digits
}

func findNthDigit(n int) int {
	d := 1 + sort.Search(8, func(length int) bool {
		return totalDigits(length+1) >= n
	})
	prevDigits := totalDigits(d - 1)
	index := n - prevDigits - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}
