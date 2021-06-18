package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func getTitleFromDetail(name string) string {
	rd, err := os.Open(filepath.Join(name))
	panicErr(err)
	defer rd.Close()
	reader := bufio.NewReader(rd)
	title := ""
	for {
		line, _, err := reader.ReadLine()
		panicErr(err, name)
		if strings.Index(string(line), "###") != -1 {
			title = string(line)
			title = strings.Trim(title, "###")
			title = strings.TrimSpace(title)
			return title
		}
	}
}

func getTitle(name string) string {
	dir := filepath.Dir(name)

	detail := filepath.Join(dir, "README.md")
	_, err := os.Stat(detail)
	if err == nil {
		return getTitleFromDetail(detail)
	}

	rd, err := os.Open(filepath.Join(name))
	panicErr(err)
	defer rd.Close()
	reader := bufio.NewReader(rd)
	title := ""
	for {
		line, _, err := reader.ReadLine()
		panicErr(err, name)
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
