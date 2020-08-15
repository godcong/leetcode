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
	title := getTitle(s)
	_, _ = writer.WriteString("| " + strings.Join(index, "."))
	_, _ = writer.WriteString(" | " + title)
	_, _ = writer.WriteString(" | " + codeName)
	_, _ = writer.WriteString(" | " + fmt.Sprintf("[GO](%v)", root+s+".go"))
	_, _ = writer.WriteString(" | " + fmt.Sprintf("[TEST](%v)", root+s+"_test.go") + " |" + "  \n")
}
