package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

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
	fmt.Println("Cookie Loaded")

	codeName := ""
	if len(os.Args) > 1 {
		codeName = os.Args[1]
	}
	var code *query.Code
	parseInt, err := strconv.ParseInt(codeName, 10, 32)
	if err != nil {
		code, err = query.GetCode(string(file), codeName)
		if err != nil {
			fmt.Println("GetCode error", err)
			return
		}
	} else {
		code, err = query.GetNumberCode(string(file), int(parseInt))
		if err != nil {
			fmt.Println("GetNumberCode error", err)
			return
		}
		for _, question := range code.Data.ProblemSetQuestionList.Questions {
			fmt.Printf("question list:%v,slug:%v\n", question.FrontendQuestionID, question.QuestionTitleSlug)
		}
	}

	name := query.WorkspaceName(code)
	fmt.Println("Workspace:", name)
	if err := query.GenCodeWorkspace(name, code); err != nil {
		fmt.Println("gen workspace error", err)
		return
	}
	path := query.GetWorkPath(name)
	//fmt.Println("content:", code.Data.Question.TranslatedContent)

	if err := query.WriteMarkdownTo(filepath.Join(path, "README.md"), code); err != nil {
		fmt.Println("write markdown error", err)
		return
	}
	fmt.Println("Code Generated:", code.Data.Question.QuestionFrontendID)

}
