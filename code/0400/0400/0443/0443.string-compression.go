package _0443

import (
	"strconv"
)

func compress(chars []byte) int {
	p1, p2, num, c := 0, 1, 1, chars[0]
	for p2 = 1; p2 <= len(chars); p2++ {
		if p2 == len(chars) || chars[p2] != c {
			chars[p1] = c
			p1++
			if num > 1 {
				ns := []byte(strconv.Itoa(num))
				for i := 0; i < len(ns); i++ {
					chars[p1] = ns[i]
					p1++
				}
			}
			num = 1
			if p2 < len(chars) {
				c = chars[p2]
			}

		} else {
			num++
		}
	}
	return p1
}
