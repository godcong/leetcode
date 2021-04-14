package other

var col = byte('A')

func excel(number int) string {
	var ret []byte
	for v := number % 26; v > 0; v = number % 26 {
		number /= 26
		ret = append([]byte{col + byte(v-1)}, ret...)
	}
	return string(ret)
}
