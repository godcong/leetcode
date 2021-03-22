package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var root = "https://github.com/godcong/leetcode/blob/master/"

func templeWrite(writer *bufio.Writer, s string) {
	name := strings.Split(onlyName(s), ".")
	index := name[:len(name)-1]
	codeName := name[len(name)-1]
	title := getTitle(s)
	_, _ = writer.WriteString("| " + strings.Join(index, "."))
	_, _ = writer.WriteString(" | " + title)
	_, _ = writer.WriteString(" | " + codeName)
	_, _ = writer.WriteString(" | " + fmt.Sprintf("[GO](%v)", root+s+".go"))
	if isExist(s + "_test.go") {
		_, _ = writer.WriteString(" | " + fmt.Sprintf("[TEST](%v)", root+s+"_test.go") + " |" + "  \n")
	} else {
		_, _ = writer.WriteString(" | " + "TEST" + " |" + "  \n")
	}

}

func isExist(name string) bool {
	_, err := os.Stat(filepath.Join(getCurrentPath(), name))
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}
