package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

const footer = "Last Commit Date: %v\n"

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
	fmt.Println("path:", filepath.Join(getCurrentPath(), "cmd", "readme", "README.md.tmpl"))
	rd, err := os.Open(filepath.Join(getCurrentPath(), "cmd", "readme", "README.md.tmpl"))
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

	rd, err = os.OpenFile(
		filepath.Join(getCurrentPath(), "README.md"),
		os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_TRUNC,
		0755,
	)
	defer rd.Close()
	panicErr(err)
	bw := bufio.NewWriter(rd)

	for _, bytes := range headLine {
		_, _ = bw.Write(bytes)
		_, _ = bw.WriteString("\n")
	}

	_, _ = bw.WriteString("### 总完成:" + strconv.Itoa(len(files)) + " ###" + "  \n")
	_, _ = bw.WriteString(
		"| 目录     |  标题                                                   |  题目名称                                                   |  函数入口 |  测试代码 |  \n")
	_, _ = bw.WriteString("|:--------:|:--------------------------------------------------------|:--------------------------------------------------------|:--------:|:--------:|  \n")
	for _, file := range files {
		templeWrite(bw, file)
	}
	_, _ = bw.WriteString("<!--END-->\n")
	if err := finishFooter(bw); err != nil {
		panicErr(err)
	}
	err = bw.Flush()
	panicErr(err)
}

func finishFooter(w io.Writer) (err error) {
	f := fmt.Sprintf(footer, time.Now().Format("2006-01-02"))
	_, err = w.Write([]byte(f))
	return err
}

func printLineArray(arr ...string) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func panicErr(err error, args ...interface{}) {
	if err != nil {
		if args != nil {
			fmt.Printf("error:%v\n", args)
		}
		debug.PrintStack()
		os.Exit(0)
	}
}

var filterList = []string{
	"README.md",
	"LICENSE",
	"def.go",
	"common",
	"common.go",
	"common_func.go",
	"go.mod",
	"go.sum",
	"_test.go",
	".gitignore",
}

func getAllFiles(path string, filters []string, filterTest bool) []string {
	dir, err := ioutil.ReadDir(path)
	panicErr(err)
	var ret []string
	var need bool
	for _, info := range dir {
		need = true
		if info.IsDir() {
			for _, filter := range filters {
				if info.Name() == filter {
					need = false
					break
				}
			}
			sub := filepath.Join(path, info.Name())
			if need {
				target, err := ioutil.ReadDir(sub)
				if err != nil {
					continue
				}
				for _, fileInfo := range target {
					name := readFile(filepath.Join(sub, fileInfo.Name()), filters, filterTest)
					if name != "" {
						ret = append(ret, name)
					}
				}
			}
		}
		for _, filter := range filters {
			if info.Name() == filter {
				need = false
				break
			}
		}
		if need {
			name := readFile(filepath.Join(path, info.Name()), filters, filterTest)
			if name != "" {
				ret = append(ret, name)
			}
		}
	}
	return ret
}

func readFile(path string, filters []string, filterTest bool) string {
	stat, err := os.Stat(path)
	if err != nil {
		return ""
	}
	if stat.IsDir() {
		return ""
	}
	//if filterTest && strings.Index(path, "_test.go") > 0 {
	//	return ""
	//}

	for _, filter := range filters {
		if strings.Contains(stat.Name(), filter) {
			return ""
		}
	}
	return path
}

func onlyName(s string) string {
	s = filepath.Base(s)
	i := strings.LastIndex(s, ".")
	return s[:i]
}
