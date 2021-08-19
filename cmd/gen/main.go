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
	cookie, err := os.ReadFile(filepath.Join(wd, "cookie"))
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
	code, err = query.GetCode(string(cookie), codeName)
	if err != nil {
		fmt.Println("GetCode error", err)
		return
	}

	name := fmt.Sprintf("%04v", code.Result.Number)

	name = query.WorkspaceName(name)
	if query.DEBUG {
		fmt.Println("Workspace:", name)
	}
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
	fmt.Println("Code Generated:", name)

}
