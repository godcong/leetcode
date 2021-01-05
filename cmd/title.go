package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func getTitle(name string) string {
	rd, err := os.Open(filepath.Join(codePath, name+".go"))
	panicErr(err)
	defer rd.Close()
	reader := bufio.NewReader(rd)
	title := ""
	for {
		line, _, err := reader.ReadLine()
		panicErr(err)
		if strings.Index(string(line), "/*") != -1 {
			line, _, err = reader.ReadLine()
			panicErr(err)
			title = string(line)
			break
		}
	}
	index := strings.LastIndex(title, ".")
	if index != -1 {
		title = title[index+1:]
	}
	return strings.TrimSpace(title)
}
