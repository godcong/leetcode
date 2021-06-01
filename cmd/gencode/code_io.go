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

	codePath := filepath.Join(wd, "code", name)

	_, err = os.Stat(codePath)
	if err == nil {
		fmt.Println(codePath, " is already exist")
		os.Exit(0)
	}
	return codePath
}

func genCodeDir(name string) {

	if name == "" {
		return
	}

	codePath := getWorkPath(name)

	_ = os.MkdirAll(codePath, 0755)

	err := ioutil.WriteFile(filepath.Join(codePath, name+".go"), packageHeader(name), 0755)
	if err != nil {
		panic(err)
	}
}

func writeCodeJSON(code Code) {
	name := code.Data.Question.Questionid
	indent, err := json.MarshalIndent(code, "", " ")
	if err != nil {
		return
	}
	codePath := getWorkPath(name)
	err = ioutil.WriteFile(filepath.Join(codePath, name+".json"), indent, 0755)
	if err != nil {
		panic(err)
	}

}

func packageHeader(name string) []byte {
	return []byte(fmt.Sprintf("package _%v", name))
}
