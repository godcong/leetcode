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
	i, j := 0, 0
	iSize, jSize := len(num1), len(num2)
	nSize := iSize + jSize
	num := make([]byte, nSize)
	for i = 1; i <= iSize; i++ {
		for j = 1; j <= jSize; j++ {
			byteAdd(&num, nSize-(i+j)+1, (num1[iSize-i]-'0')*(num2[jSize-j]-'0'))
		}
	}

	if num[0] == 0 {
		num = num[1:]
	}

	for i = range num {
		(num)[i] += '0'
	}

	return string(num)
}

func byteAdd(num *[]byte, index int, b byte) {
	for b > 0 {
		b, (*num)[index], index = (b+(*num)[index])/10, (b+(*num)[index])%10, index-1
	}
}
