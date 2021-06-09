package main

import (
	"fmt"
	"path/filepath"

	"github.com/godcong/leetcode/query"
)

func main() {
	code, err := query.GetCode()
	if err != nil {
		return
	}

	if err := query.GenCodeWorkspace(code); err != nil {
		return
	}
	path := query.GetWorkPath(code.Data.Question.QuestionFrontendID)

	if err := query.WriteMarkdownTo(filepath.Join(path, "README.md"), code); err != nil {
		return
	}
	fmt.Println("Code Generated:", code.Data.Question.QuestionFrontendID)

}
