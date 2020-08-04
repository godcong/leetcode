package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ret := 0
	abs := false
	if x < 0 {
		abs = true
		x = -x
	}
	for x > 9 {
		ret += x % 10
		x /= 10
		ret *= 10
	}
	ret += x
	if abs {
		return -ret
	}
	if ret > math.MaxInt32 {
		return 0
	}
	return ret
}

func main() {
	fmt.Println("123", reverse(123))
	fmt.Println("-123", reverse(-123))
	fmt.Println("120", reverse(120))
	fmt.Println("1534236469", reverse(1534236469))
}
