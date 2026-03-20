package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/godcong/leetcode/internal/markdown"
	"github.com/godcong/leetcode/internal/types"
)

var codeGroups = map[string]string{
	//"default": "code",
	"SwordRefers":       "offer",   // Legacy: Sword Refers Offer
	"InterviewQuestion": "qustion", // Legacy: Interview Questions
	"LCR":               "lcr",     // New: LeetCode Classic Review
	"LCP":               "lcp",     // New: LeetCode Problems (Contest)
}

var replaceChinese = map[string]string{
	"面试题": "InterviewQuestion",
	"剑指":   "SwordRefers",
	".":      "_",
	"-":      "_",
	" ":      "_",
}

var workspaceGroups = []string{
	"InterviewQuestion",
	"SwordRefers_Offer_I",
	"SwordRefers_Offer_II",
}

// GetWorkPath returns the full path for a problem directory
func GetWorkPath(group string, name string) string {
	wd, err := os.Getwd()
	if err != nil {
		return "code"
	}

	basePath := filepath.Join(wd, "code")

	// Determine category and version based on name prefix
	category := "problems" // default
	version := ""

	for s, s2 := range codeGroups {
		if strings.Contains(name, s) {
			// Legacy rules: SwordRefers -> offer, InterviewQuestion -> qustion
			if s2 == "offer" || s2 == "qustion" {
				category = s2
			} else if s2 == "lcr" || s2 == "lcp" {
				// New rules: LCR/LCP -> problems/lcr or problems/lcp
				category = fmt.Sprintf("problems/%s", s2)
			}
			version = s2
			break
		}
	}

	// Build directory structure based on type
	var path string
	if version != "" && (version == "offer" || version == "qustion") {
		// Legacy: offer/{range}/{name}/ or qustion/{range}/{name}/
		rangeDir := getRangeGroup(name)
		path = filepath.Join(basePath, category, rangeDir, name)
	} else if version != "" && (version == "lcr" || version == "lcp") {
		// New: problems/lcr/{range}/{name}/ or problems/lcp/{range}/{name}/
		rangeDir := getRangeGroup(name)
		path = filepath.Join(basePath, category, rangeDir, name)
	} else if IsNumber(name) {
		// Normal numbers: problems/{range}/{number}/
		rangeDir := getRangeGroup(name)
		path = filepath.Join(basePath, category, rangeDir, name)
	} else {
		// Other slugs: problems/{name}/
		path = filepath.Join(basePath, category, name)
	}

	_ = os.MkdirAll(path, 0755)
	return path
}

// getRangeGroup returns the range directory (e.g., "0000-0099", "0100-0199")
func getRangeGroup(name string) string {
	// Extract number from name - find all consecutive digits
	numStr := ""
	for _, ch := range name {
		if ch >= '0' && ch <= '9' {
			numStr += string(ch)
		} else if numStr != "" {
			// Stop at first non-digit after we've started collecting digits
			break
		}
	}

	if numStr == "" {
		return "nogroup"
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "nogroup"
	}

	// Calculate range
	rangeStart := (num / 100) * 100
	rangeEnd := rangeStart + 99

	return fmt.Sprintf("%04d-%04d", rangeStart, rangeEnd)
}

func getNum(num string) string {
	nums := strings.Split(num, " ")
	if len(nums) == 0 {
		return num
	}
	return nums[len(nums)-1]
}

// GetGroupName returns the group name for a problem number
func GetGroupName(num string) string {
	nnum := getNum(num)
	for _, group := range workspaceGroups {
		if strings.Contains(nnum, group) {
			return group
		}
	}

	if IsNumber(nnum) {
		if len(nnum) > 2 {
			return strings.ReplaceAll(num, nnum, fmt.Sprintf("%02v00", nnum[0:len(nnum)-2]))
		}
		return "0000"
	}

	return "nogroup"
}

// GenCodeWithPath generates the Go solution file
func GenCodeWithPath(path string, name string, code *types.Code) (string, error) {
	codeGo := filepath.Join(path, fmt.Sprintf("%v.go", code.Result.Slug))
	_, err := os.Stat(codeGo)
	if err == nil {
		return codeGo, nil
	}

	file, err := os.OpenFile(codeGo, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_SYNC, 0755)
	if err != nil {
		return codeGo, fmt.Errorf("open file:%v", err)
	}
	defer file.Close()
	file.Write(packageHeader(name))
	file.WriteString("\n")
	file.Write(importCommon())
	file.WriteString("\n")
	for i := range code.Data.Question.CodeSnippets {
		if code.Data.Question.CodeSnippets[i].Lang == "Go" {
			file.WriteString(code.Data.Question.CodeSnippets[i].Code)

		}
	}
	return codeGo, nil
}

// GenerateREADME generates README.md for a problem with optional date
func GenerateREADME(path string, code *types.Code, date string) error {
	readmePath := filepath.Join(path, "README.md")
	return markdown.WriteMarkdownTo(readmePath, code, date)
}

// WriteCodeJSON writes code data to JSON file
func WriteCodeJSON(name string, code types.Code) {
	indent, err := json.MarshalIndent(code, "", " ")
	if err != nil {
		return
	}
	codePath := GetWorkPath(name, "")
	_ = os.MkdirAll(codePath, 0755)
	err = ioutil.WriteFile(filepath.Join(codePath, name+".json"), indent, 0755)
	if err != nil {
		panic(err)
	}
}

func packageHeader(name string) []byte {
	return []byte(fmt.Sprintf("package _%v\n", name))
}

func importCommon() []byte {
	return []byte(fmt.Sprintf("import (\n\t. \"github.com/godcong/leetcode/common\"\n)\n"))
}

// GetNameCleanupChinese cleans up Chinese characters from problem names
func GetNameCleanupChinese(name string) string {
	if name == "" {
		fmt.Println("empty name error")
		return ""
	}

	for k, v := range replaceChinese {
		name = strings.ReplaceAll(name, k, v)
	}

	for strings.Contains(name, "__") {
		name = strings.ReplaceAll(name, "__", "_")
	}

	if _, err := strconv.Atoi(name); err == nil {
		return fmt.Sprintf("%04v", name)
	}
	return name
}

// IsNumber checks if a string is a number
func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
