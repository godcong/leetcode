package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

	if err := addToGit(path, code.Result.Slug); err != nil {
		fmt.Println("add to git error", err)
		return
	}

}

func addToGit(path string, name string) error {
	command := exec.Command("git", "version")
	ret, err := command.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(ret))
	if !strings.Contains(string(ret), "git version") {
		//skip if git is not exist
		return nil
	}

	//command = exec.Command("git", "add", path)
	//_, err = command.CombinedOutput()
	//if err != nil {
	//	return fmt.Errorf("add to git error: %v", err)
	//}

	command = exec.Command("git", "commit", "-a", "-m", fmt.Sprintf("add %s", name))
	_, err = command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("commit to git error: %v", err)
	}

	return nil
}
