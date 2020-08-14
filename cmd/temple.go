package main

import (
	"bufio"
	"fmt"
	"strings"
)

var root = "https://github.com/godcong/leetcode/blob/master/"

func templeWrite(writer *bufio.Writer, s string) {
	name := strings.Split(s, ".")
	index := name[:len(name)-1]
	codeName := name[len(name)-1]

	_, _ = writer.WriteString("## 索引号: " + strings.Join(index, ".") + " ##  \n")
	_, _ = writer.WriteString("实现函数: " + codeName + "  \n")
	_, _ = writer.WriteString("实现代码: " + fmt.Sprintf("[%v](%v)", s+".go", root+s+".go  \n"))
	_, _ = writer.WriteString("测试代码: " + fmt.Sprintf("[%v](%v)", s+"_test.go", root+s+"_test.go  \n"))
	_, _ = writer.WriteString("\n")
}
