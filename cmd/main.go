package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var headLine [][]byte
var codePath = filepath.Join(getCurrentPath(), "code")

func init() {

}

func getCurrentPath() string {
	getwd, err := os.Getwd()
	panicErr(err)
	return getwd
}

//create readme
func main() {
	fmt.Println("path:", filepath.Join(getCurrentPath(), "README.md"))
	rd, err := os.Open(filepath.Join(getCurrentPath(), "README.md"))
	panicErr(err)
	reader := bufio.NewReader(rd)
	for {
		line, _, err := reader.ReadLine()
		fmt.Println("line:", string(line))
		panicErr(err)
		headLine = append(headLine, line)
		if strings.Compare("<!--STA-->", string(line)) == 0 {
			break
		}
	}
	err = rd.Close()
	panicErr(err)

	files := getAllFiles(codePath, filterList, true)
	printLineArray(files...)

	rd, err = os.OpenFile(filepath.Join(getCurrentPath(), "README.md"), os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_TRUNC, 0755)
	defer rd.Close()
	panicErr(err)
	bw := bufio.NewWriter(rd)

	for _, bytes := range headLine {
		_, _ = bw.Write(bytes)
		_, _ = bw.WriteString("\n")
	}

	_, _ = bw.WriteString("### 总完成:" + strconv.Itoa(len(files)) + " ###" + "  \n")
	_, _ = bw.WriteString("| 目录     |  标题                                                   |  函数名                                                   |  实现代码 |  测试代码 |  \n")
	_, _ = bw.WriteString("|:--------:|:--------------------------------------------------------|:--------------------------------------------------------|:--------:|:--------:|  \n")
	for _, file := range files {
		templeWrite(bw, file)
	}
	_, _ = bw.WriteString("<!--END-->\n")

	err = bw.Flush()
	panicErr(err)
}

func printLineArray(arr ...string) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

var filterList = []string{
	"README.md",
	"LICENSE",
	"def.go",
	"common.go",
	"common_func.go",
	"go.mod",
	"go.sum",
	".gitignore",
}

func getAllFiles(path string, filters []string, filterTest bool) []string {
	dir, err := ioutil.ReadDir(path)
	panicErr(err)
	need := false
	var ret []string
	for _, info := range dir {
		need = true
		if info.IsDir() {
			continue
		}
		if filterTest && strings.Index(info.Name(), "_test.go") > 0 {
			continue
		}

		for _, filter := range filters {
			if info.Name() == filter {
				need = false
			}
		}
		if need {
			ret = append(ret, onlyName(info.Name()))
		}
	}
	return ret
}

func onlyName(s string) string {
	i := strings.LastIndex(s, ".")
	return s[:i]
}
