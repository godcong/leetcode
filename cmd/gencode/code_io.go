package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getWorkPath(name string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd, "code", name)
}

func genCodeDir(name string, code Code) {

	if name == "" {
		return
	}

	codePath := getWorkPath(name)

	_ = os.MkdirAll(codePath, 0755)

	file, err := os.OpenFile(filepath.Join(codePath, name+".go"), os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0755)
	if err != nil {
		panic(err)
	}
	file.Write(packageHeader(name))
	file.WriteString("\n")
	for i := range code.Data.Question.Codesnippets {
		if code.Data.Question.Codesnippets[i].Lang == "Go" {
			file.WriteString(code.Data.Question.Codesnippets[i].Code)
		}
	}

}

func writeCodeJSON(name string, code Code) {
	indent, err := json.MarshalIndent(code, "", " ")
	if err != nil {
		return
	}
	codePath := getWorkPath(name)
	_ = os.MkdirAll(codePath, 0755)
	err = ioutil.WriteFile(filepath.Join(codePath, name+".json"), indent, 0755)
	if err != nil {
		panic(err)
	}
}

func packageHeader(name string) []byte {
	return []byte(fmt.Sprintf("package _%v", name))
}
