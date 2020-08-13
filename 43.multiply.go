package leetcode

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num := new([]byte)
	i, j := 0, 0
	iSize := len(num1) - 1
	jSize := len(num2) - 1
	m := byte(0)
	for i = 0; i <= iSize; i++ {
		for j = 0; j <= jSize; j++ {
			m = (num1[iSize-i] - '0') * (num2[jSize-j] - '0')
			byteDec(num, i+j, m%10)
			if m/10 > 0 {
				byteDec(num, i+j+1, m/10)
			}
		}
	}
	return string(*num)
}

func byteDec(num *[]byte, index int, b byte) {
	nidx := len(*num) - (index) - 1
	if nidx < 0 {
		*num = []byte(string(b+'0') + string(*num))
		return
	}
	m := ((*num)[nidx] - '0') + (b)
	(*num)[nidx] = m%10 + '0'
	if m > 9 {
		byteDec(num, index+1, m/10)
	}
}
