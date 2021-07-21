package query

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var pathPaths = map[string]string{
	//"default": "code",
	"SwordRefers":       "offer",
	"InterviewQuestion": "qustion",
}

func GetWorkPath(name string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := "code"
	for s, s2 := range pathPaths {
		if strings.Contains(name, s) {
			path = s2
		}
	}

	return filepath.Join(wd, path, name)
}

var replaceList = map[string]string{
	"面试题": "InterviewQuestion",
	"剑指":  "SwordRefers",
	".":   "_",
	"-":   "_",
}

func GenCodeWorkspace(name string, code *Code) error {

	path := GetWorkPath(name)

	_ = os.MkdirAll(path, 0755)

	codeGo := filepath.Join(path, fmt.Sprintf("%v.%v.go", name, code.Data.Question.TitleSlug))
	_, err := os.Stat(codeGo)
	if err == nil {
		return nil
	}

	file, err := os.OpenFile(codeGo, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_SYNC, 0755)
	if err != nil {
		return fmt.Errorf("open file:%v", err)
	}
	defer file.Close()
	file.Write(packageHeader(name))
	file.WriteString("\n")
	file.WriteString("\n")
	for i := range code.Data.Question.CodeSnippets {
		if code.Data.Question.CodeSnippets[i].Lang == "Go" {
			file.WriteString(code.Data.Question.CodeSnippets[i].Code)
		}
	}
	return nil
}

func WriteCodeJSON(name string, code Code) {
	indent, err := json.MarshalIndent(code, "", " ")
	if err != nil {
		return
	}
	codePath := GetWorkPath(name)
	_ = os.MkdirAll(codePath, 0755)
	err = ioutil.WriteFile(filepath.Join(codePath, name+".json"), indent, 0755)
	if err != nil {
		panic(err)
	}
}

func packageHeader(name string) []byte {
	return []byte(fmt.Sprintf("package _%v", name))
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

func WorkspaceName(code *Code) string {
	name := fmt.Sprintf("%04v", code.Data.Question.QuestionFrontendID)

	if name == "" {
		fmt.Println("empty name error")
		return ""
	}

	name = strings.ReplaceAll(name, " ", "_")
	for k, v := range replaceList {
		name = strings.ReplaceAll(name, k, v)
	}
	name = strings.ReplaceAll(name, "_-_", "_")
	name = strings.ReplaceAll(name, "___", "_")

	fmt.Println("Generate name:", name)

	return name
}
