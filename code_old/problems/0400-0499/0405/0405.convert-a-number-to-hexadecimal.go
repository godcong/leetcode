package _0405

import (
	"strings"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	builder := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (4 * i) & 0xf
		if val > 0 || builder.Len() > 0 {
			var digit byte
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			builder.WriteByte(digit)
		}
	}
	return builder.String()
}