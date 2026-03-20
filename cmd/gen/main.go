package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/godcong/leetcode/internal/fetcher"
	"github.com/godcong/leetcode/internal/git"
	"github.com/godcong/leetcode/internal/generator"
	"github.com/godcong/leetcode/internal/logger"
	"github.com/godcong/leetcode/internal/types"
)

func main() {
	// Parse flags
	debug := flag.Bool("debug", false, "Enable debug logging with structured output")
	noGit := flag.Bool("no-git", false, "Disable git operations (for debugging)")
	flag.Parse()
	
	// Initialize logger
	logger.Init(*debug)
	
	logger.Info("Starting gen tool", "args", os.Args[1:])
	
	// Check git only if not disabled
	if !*noGit {
		if err := git.Check(); err != nil {
			logger.Error("Git check failed", err)
			return
		}
		logger.Verbose("Git check passed")
	} else {
		logger.Info("Git operations disabled")
	}

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Failed to get working directory", err)
		return
	}
	cookie, err := os.ReadFile(filepath.Join(wd, "cookie"))
	if err != nil {
		logger.Error("Failed to read cookie", err)
		return
	}

	codeName := ""
	if len(os.Args) > 1 {
		// Skip flag arguments
		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
				codeName = arg
				break
			}
		}
	}
	
	// If no codeName provided, fetch today's daily question
	if codeName == "" {
		logger.Info("No problem specified, fetching today's daily question")
		q := fetcher.NewQuery(string(cookie))
		
		// Fetch daily question records using public method
		dailyRecords, err := q.DailyQuestionRecords(time.Now())
		if err != nil {
			logger.Error("Failed to fetch daily question records", err)
			return
		}
		
		if len(dailyRecords.Data.DailyQuestionRecords) == 0 {
			logger.Error("No daily question found for today", fmt.Errorf("empty daily records"))
			return
		}
		
		// Get today's question
		todayRecord := dailyRecords.Data.DailyQuestionRecords[0]
		codeName = todayRecord.Question.QuestionTitleSlug
		logger.Info("Today's daily question", "slug", codeName, "date", todayRecord.Date)
	}
	
	logger.Info("Fetching problem", "name", codeName)
	
	// Create fetcher and code instances
	q := fetcher.NewQuery(string(cookie))
	code := &types.Code{}
	
	// Fetch problem data
	logger.Debug("Calling GetCode", "name", codeName)
	if err := q.GetCode(code, codeName); err != nil {
		logger.Error("GetCode failed", err)
		return
	}
	
	// Validate fetched data
	if code.Data.Question.Title == "" || code.Result.Slug == "" {
		logger.Error("Invalid problem data: missing title or slug", 
			fmt.Errorf("empty data"),
			"title", code.Data.Question.Title,
			"slug", code.Result.Slug)
		return
	}
	
	logger.Info("Fetched problem",
		"number", code.Result.Number,
		"title", code.Data.Question.Title)

	name := generator.GetNameCleanupChinese(code.Result.Number)
	group := generator.GetGroupName(name)

	logger.Verbose("Processing problem info", "originalName", code.Result.Number, "cleanedName", name, "group", group)

	path := generator.GetWorkPath(group, name)

	logger.Verbose("Generated paths", "path", path)
	logger.Info("Generating files for", "number", code.Result.Number, "title", code.Data.Question.Title)

	readmePath := filepath.Join(path, "README.md")
	_, err = os.Stat(readmePath)
	if err == nil {
		logger.Info("README already exists", "path", readmePath)
		existFile := filepath.Join(path, time.Now().Format("20060102.exist"))
		_, err := os.Create(existFile)
		if err != nil {
			logger.Error("Failed to create exist file", err, "path", existFile)
			logger.Error("Skipping existing file", err)
		}
		if err := git.AddAndCommit(existFile, fmt.Sprintf("code(%v) exists", name), !*noGit); err != nil {
			logger.Error("Failed to add to git", err, "file", existFile)
			return
		}
		logger.Info("Path already exists", "path", path)
		return
	}

	// Generate README with date if available
	date := ""
	if len(code.Data.DailyQuestionRecords) > 0 {
		date = code.Data.DailyQuestionRecords[0].Date
	} else if len(code.Data.TodayRecord) > 0 {
		date = code.Data.TodayRecord[0].Date
	}
	
	if err := generator.GenerateREADME(path, code, date); err != nil {
		logger.Error("Failed to write README", err)
		return
	}
	logger.Verbose("README written", "path", readmePath, "date", date)
	if err := git.AddAndCommit(readmePath, fmt.Sprintf("code(%v) readme", name), !*noGit); err != nil {
		logger.Error("Failed to add README to git", err)
		return
	}

	codePath, err := generator.GenCodeWithPath(path, name, code)
	if err != nil {
		logger.Error("Failed to generate Go file", err)
		return
	}
	logger.Verbose("Go file generated", "path", codePath)
	if err := git.AddAndCommit(codePath, fmt.Sprintf("code(%v) %v", name, code.Result.Slug), !*noGit); err != nil {
		logger.Error("Failed to add Go file to git", err)
		return
	}

	testPath, err := createTestFile(codePath)
	if err != nil {
		logger.Error("Failed to create test file", err)
		return
	}
	logger.Verbose("Test file created", "path", testPath)
	if err := git.AddAndCommit(testPath, fmt.Sprintf("code(%v) test for:%v", name, code.Result.Slug), !*noGit); err != nil {
		logger.Error("Failed to add test file to git", err)
		return
	}

	logger.Info("Generated", 
		"number", code.Result.Number,
		"title", code.Data.Question.Title,
		"path", path)
}

func createTestFile(path string) (string, error) {
	command := exec.Command("gotests")
	_, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}

	dir, name := filepath.Split(path)
	testPath := filepath.Join(dir, strings.TrimSuffix(name, ".go")+"_test.go")
	command = exec.Command("gotests", "--all", "-w", testPath, path)
	_, err = command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("gotests create error: %v", err)
	}

	return testPath, nil
}
