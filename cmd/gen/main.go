package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/godcong/leetcode/query"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	file, err := os.ReadFile(filepath.Join(wd, "cookie"))
	if err != nil {
		return
	}

	code, err := query.GetCode(string(file))
	if err != nil {
		return
	}

	name := fmt.Sprintf("%04v", code.Data.Question.QuestionFrontendID)

	if err := query.GenCodeWorkspace(name, code); err != nil {
		return
	}
	path := query.GetWorkPath(name)

	if err := query.WriteMarkdownTo(filepath.Join(path, "README.md"), code); err != nil {
		return
	}
	fmt.Println("Code Generated:", code.Data.Question.QuestionFrontendID)

}
