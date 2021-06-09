package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetWorkPath(name string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd, "code", name)
}

func GenCodeWorkspace(name string, code *Code) error {
	if name == "" {
		return errors.New("empty name")
	}

	codePath := GetWorkPath(name)

	_ = os.MkdirAll(codePath, 0755)

	codeGo := filepath.Join(codePath, name+".go")
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
