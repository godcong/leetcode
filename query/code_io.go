package query

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var codeGroups = map[string]string{
	//"default": "code",
	"SwordRefers":       "offer",
	"InterviewQuestion": "qustion",
}

var replaceChinese = map[string]string{
	"面试题": "InterviewQuestion",
	"剑指":   "SwordRefers",
	".":      "_",
	"-":      "_",
	" ":      "_",
}

var workspaceGroups = []string{
	"InterviewQuestion",
	"SwordRefers_Offer_I",
	"SwordRefers_Offer_II",
}

// GetWorkPath ...
// @Description: 获取工作目录
// @param string
// @return string
func GetWorkPath(group string, name string) string {
	if DEBUG {
		fmt.Println("GetWorkPath", "group", group, "name", name)
	}
	path := "code"
	wd, err := os.Getwd()
	if err != nil {
		return path
	}
	for s, s2 := range codeGroups {
		if strings.Contains(name, s) {
			path = s2
		}
	}

	path = filepath.Join(wd, path, group, name)
	_ = os.MkdirAll(path, 0755)
	return path
}

func getNum(num string) string {
	nums := strings.Split(num, " ")
	if len(nums) == 0 {
		return num
	}
	return nums[len(nums)-1]
}

// GetGroupName ...
// @Description: 获取组名
func GetGroupName(num string) string {
	if DEBUG {
		fmt.Println("GroupName: ", "number", num)
	}

	nnum := getNum(num)
	for k, v := range replaceChinese {
		num = strings.ReplaceAll(num, k, v)
	}

	for _, group := range workspaceGroups {
		if strings.Contains(nnum, group) {
			return group
		}
	}

	if IsNumber(nnum) {
		if len(nnum) > 2 {
			return strings.ReplaceAll(num, nnum, fmt.Sprintf("%02v00", nnum[0:len(nnum)-2]))
		}
		return "0000"
	}

	return "nogroup"
}

func GenCodeWorkspace(path string, name string, code *Code) (string, error) {
	codeGo := filepath.Join(path, fmt.Sprintf("%v.%v.go", name, code.Result.Slug))
	_, err := os.Stat(codeGo)
	if err == nil {
		return codeGo, nil
	}

	file, err := os.OpenFile(codeGo, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_SYNC, 0755)
	if err != nil {
		return codeGo, fmt.Errorf("open file:%v", err)
	}
	defer file.Close()
	file.Write(packageHeader(name))
	file.WriteString("\n")
	file.Write(importCommon())
	file.WriteString("\n")
	for i := range code.Data.Question.CodeSnippets {
		if DEBUG {
			fmt.Printf("question code:%+v", code.Data.Question.CodeSnippets[i])
		}

		if code.Data.Question.CodeSnippets[i].Lang == "Go" {
			file.WriteString(code.Data.Question.CodeSnippets[i].Code)

		}
	}
	return codeGo, nil
}

func WriteCodeJSON(name string, code Code) {
	indent, err := json.MarshalIndent(code, "", " ")
	if err != nil {
		return
	}
	codePath := GetWorkPath(name, "")
	_ = os.MkdirAll(codePath, 0755)
	err = ioutil.WriteFile(filepath.Join(codePath, name+".json"), indent, 0755)
	if err != nil {
		panic(err)
	}
}

func packageHeader(name string) []byte {
	return []byte(fmt.Sprintf("package _%v\n", name))
}

func importCommon() []byte {
	return []byte(fmt.Sprintf("import (\n\t. \"github.com/godcong/leetcode/common\"\n)\n"))
}

func decodeCode(closer io.ReadCloser, code *Code) error {
	all, err := ioutil.ReadAll(closer)
	if err != nil {
		return err
	}
	//fmt.Println("decode", string(all))

	//ioutil.WriteFile("tmp.txt", all, 0755)
	return json.Unmarshal(all, code)
}

//GetNameCleanupChinese ...
func GetNameCleanupChinese(name string) string {
	if name == "" {
		fmt.Println("empty name error")
		return ""
	}

	for k, v := range replaceChinese {
		name = strings.ReplaceAll(name, k, v)
	}

	for strings.Contains(name, "__") {
		name = strings.ReplaceAll(name, "__", "_")
	}
	if DEBUG {
		fmt.Println("Generate name:", name)
	}

	if _, err := strconv.Atoi(name); err == nil {
		return fmt.Sprintf("%04v", name)
	}
	return name
}

func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
