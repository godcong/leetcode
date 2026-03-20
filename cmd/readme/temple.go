package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var root = ""

func templeWrite(writer *bufio.Writer, s string) {
	name := strings.Split(onlyName(s), ".")
	index := name[:len(name)-1]
	codeName := name[len(name)-1]
	title := getTitle(s)
	if DEBUG {
		fmt.Println("title:", title)
	}
	_, _ = writer.WriteString("| " + strings.Join(index, "."))
	_, _ = writer.WriteString(" | " + title)
	_, _ = writer.WriteString(" | " + codeName)
	
	// Convert to relative path with forward slashes (remove .go extension from path)
	relPath := strings.TrimPrefix(s, getCurrentPath()+string(filepath.Separator))
	relPath = filepath.ToSlash(relPath)
	// Remove trailing .go if present (we add it in the link)
	relPath = strings.TrimSuffix(relPath, ".go")
	
	_, _ = writer.WriteString(" | " + fmt.Sprintf("[GO](%v.go)", relPath))
	if isExist(s + "_test.go") {
		_, _ = writer.WriteString(" | " + fmt.Sprintf("[TEST](%v_test.go)", relPath) + " |" + "  \n")
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
