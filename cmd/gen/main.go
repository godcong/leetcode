package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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
	group := query.GetGroupName(code.Result.Number)
	if query.DEBUG {
		fmt.Println("Group:", group)
	}

	path := query.GetWorkPath(group, name)

	//fmt.Println("content:", code.Data.Question.TranslatedContent)
	fmt.Println("Code Generate:", name)
	readmePath := filepath.Join(path, "README.md")
	if err := query.WriteMarkdownTo(readmePath, code); err != nil {
		fmt.Println("write markdown error", err)
		return
	}
	if err := addToGit(readmePath, fmt.Sprintf("code(%v) readme", code.Result.Number)); err != nil {
		fmt.Println("add to git error", err)
		return
	}

	codePath, err := query.GenCodeWorkspace(path, name, code)
	if err != nil {
		fmt.Println("gen workspace error", err)
		return
	}
	if err := addToGit(codePath, fmt.Sprintf("code(%v) %v", code.Result.Number, code.Result.Slug)); err != nil {
		fmt.Println("add to git error", err)
		return
	}

	testPath, err := createTestFile(codePath)
	if err != nil {
		fmt.Println("create test file error", err)
		return
	}
	if err := addToGit(testPath, fmt.Sprintf("code(%v) test for:%v", code.Result.Number, code.Result.Slug)); err != nil {
		fmt.Println("add to git error", err)
		return
	}

	fmt.Println("Generated Time:", time.Now().Format(time.RFC3339))

}

func createTestFile(path string) (string, error) {
	command := exec.Command("gotests")
	ret, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	if query.DEBUG {
		fmt.Println("gotests found", string(ret))
	}

	dir, name := filepath.Split(path)
	testPath := filepath.Join(dir, strings.TrimSuffix(name, ".go")+"_test.go")
	command = exec.Command("gotests", "--all", "-w", testPath, path)
	ret, err = command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("gotests create error: %v", err)
	}
	if query.DEBUG {
		fmt.Println("gotests create", string(ret))
	}

	return testPath, nil
}

func addToGit(path string, name string) error {
	command := exec.Command("git", "version")
	ret, err := command.CombinedOutput()
	if err != nil {
		return err
	}
	if query.DEBUG {
		fmt.Println("git version:", string(ret))
	}

	if !strings.Contains(string(ret), "git version") {
		//skip if git is not exist
		return nil
	}

	if query.DEBUG {
		fmt.Println("add path to git:", path)
	}

	command = exec.Command("git", "add", path)
	_, err = command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("add to git error: %v", err)
	}

	command = exec.Command("git", "commit", "-a", "-m", fmt.Sprintf("feat(code): add new %s", name))
	_, err = command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("commit to git error: %v", err)
	}

	return nil
}
