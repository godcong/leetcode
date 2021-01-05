package code

import "strconv"

/*
93. 复原IP地址
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。



示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/
func restoreIpAddresses(s string) []string {
	return getAddress(s, 0, 3, "")
}

func getAddress(s string, start, deep int, prefix string) []string {
	var ret []string
	if deep == 0 {
		if s[start] == '0' && len(s[start:]) > 1 {
			return nil
		}

		if addri, err := strconv.Atoi(s[start:]); err != nil || addri > 255 {
			return nil
		}
		return []string{prefix + s[start:]}
	}
	for i := 1; start+i < len(s); i++ {
		if s[start] == '0' && len(s[start:start+i]) > 1 {
			break
		}
		if addri, err := strconv.Atoi(s[start : start+i]); err != nil || addri > 255 {
			break
		}
		ret = append(ret, getAddress(s, start+i, deep-1, prefix+s[start:start+i]+".")...)
	}
	return ret
}
