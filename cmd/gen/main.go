package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/godcong/leetcode/query"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	file, err := os.ReadFile(filepath.Join(wd, "cookie"))
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("cookie:", string(file))

	codeName := ""
	if len(os.Args) > 1 {
		codeName = os.Args[1]
	}

	code, err := query.GetCode(string(file), codeName)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	name := fmt.Sprintf("%04v", code.Data.Question.QuestionFrontendID)
	name = strings.ReplaceAll(name, " ", "_")
	if err := query.GenCodeWorkspace(name, code); err != nil {
		return
	}
	path := query.GetWorkPath(name)

	if err := query.WriteMarkdownTo(filepath.Join(path, "README.md"), code); err != nil {
		return
	}
	fmt.Println("Code Generated:", code.Data.Question.QuestionFrontendID)

}
