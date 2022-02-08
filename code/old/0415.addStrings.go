package old

func getAddStringsNum(b byte) int {
	return int(b - 48)
}

func getAddStringsSumByte(num1 string, num2 string) ([]byte, []byte) {
	if len(num1) > len(num2) {
		return []byte(num2), []byte(num1)
	}
	return []byte(num1), []byte(num2)
}

func addStringsSum(num []byte, index int, add int) (ret []byte) {
	if index < 0 {
		return append([]byte{'1'}, num...)
	}
	add = getAddStringsNum(num[index]) + add
	num[index] = byte(add%10 + 48)
	if (add) > 9 {
		num = addStringsSum(num, index-1, 1)
	}
	return num
}

/*
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。



提示：

num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
*/
func addStrings(num1 string, num2 string) string {
	n1, n2 := getAddStringsSumByte(num1, num2)
	loop := len(n1) - 1
	for i := 1; loop >= 0; loop-- {
		n2 = addStringsSum(n2, len(n2)-i, getAddStringsNum(n1[loop]))
		i++
	}
	return string(n2)
}
