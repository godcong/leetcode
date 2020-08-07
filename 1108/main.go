package main

import (
	"fmt"
)

//func defangIPaddr(address string) string {
//	return strings.ReplaceAll(address, ".", "[.]")
//}

//func defangIPaddr(address string) string {
//	return strings.Join(strings.Split(address, "."), "[.]")
//}

/*
62 / 62 个通过测试用例
状态：通过
执行用时：0 ms
内存消耗：1.9 MB
https://leetcode-cn.com/submissions/detail/95620031/
 */
func defangIPaddr(address string) string {
	ret := make([]byte, len(address)+6)
	start := 0
	j := 0
	for i := 1; i < len(address)-1; i++ {
		if address[i] == '.' {
			start += copy(ret[start:], address[j:i])
			start += copy(ret[start:], "[.]")
			j = i + 1
		}
	}
	copy(ret[start:], address[j:])
	return string(ret)
}

/*
给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。

所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。

 

示例 1：

输入：address = "1.1.1.1"
输出："1[.]1[.]1[.]1"
示例 2：

输入：address = "255.100.50.0"
输出："255[.]100[.]50[.]0"
 

提示：

给出的 address 是一个有效的 IPv4 地址

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/defanging-an-ip-address
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println("address = \"1.1.1.1\"", defangIPaddr("1.1.1.1"))
	fmt.Println("address = \"255.100.50.0\"", defangIPaddr("255.100.50.0"))
}
