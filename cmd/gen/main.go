package main

import (
	"fmt"

	"github.com/godcong/leetcode/query"
)

func main() {
	if code, err := query.GetCode(); err != nil {
		return
	}
	if err != nil {
		return
	}
	name := fmt.Sprintf("%04v", code.Data.Question.QuestionFrontendID)
	fmt.Println("gen code", name)
	genCodeDir(name, code)
	writeCodeJSON(name, code)
}
