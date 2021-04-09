package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	if name == "" {
		return
	}

	codePath := filepath.Join(wd, "code", name)

	_, err = os.Stat(codePath)
	if err == nil {
		fmt.Println(codePath, " is already exist")
		os.Exit(0)
	}

	_ = os.MkdirAll(codePath, 0755)

	err = ioutil.WriteFile(filepath.Join(codePath, name+".go"), packageHeader(name), 0755)
	if err != nil {
		panic(err)
	}
}

func packageHeader(name string) []byte {
	return []byte(fmt.Sprintf("package _%v", name))
}
