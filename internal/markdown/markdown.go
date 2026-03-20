package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	markdown "github.com/JohannesKaufmann/html-to-markdown/v2"

	"github.com/godcong/leetcode/internal/types"
)

// WriteMarkdownTo writes the problem description to a README file
func WriteMarkdownTo(path string, code *types.Code, date string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_SYNC|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write title
	title := code.Data.Question.TranslatedTitle
	if title == "" {
		title = code.Data.Question.Title
	}

	_, err = file.WriteString(fmt.Sprintf("# %s\n\n", title))
	if err != nil {
		return err
	}

	// Write problem info including date if available
	_, err = file.WriteString("## Problem Info\n\n")
	if err != nil {
		return err
	}

	// Problem number
	number := code.Result.Number
	if number != "" {
		_, err = file.WriteString(fmt.Sprintf("- **Number:** %s\n", number))
		if err != nil {
			return err
		}
	}

	// Date (for daily problems)
	if date != "" {
		_, err = file.WriteString(fmt.Sprintf("- **Daily Question Date:** %s\n", date))
		if err != nil {
			return err
		}
	}

	// Difficulty
	difficulty := code.Data.Question.Difficulty
	if difficulty != "" {
		_, err = file.WriteString(fmt.Sprintf("- **Difficulty:** %s\n", difficulty))
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	// Write problem link
	titleSlug := code.Data.Question.TitleSlug
	if titleSlug != "" {
		_, err = file.WriteString(fmt.Sprintf("## Problem Link\n\n- [LeetCode](https://leetcode-cn.com/problems/%s/)\n\n", titleSlug))
		if err != nil {
			return err
		}
	}

	// Write description
	_, err = file.WriteString("## Description\n\n")
	if err != nil {
		return err
	}

	// Preprocess content to fix common issues
	content := code.Data.Question.TranslatedContent
	if content == "" {
		content = code.Data.Question.Content
	}
	content = preprocessContent(content)

	// Convert HTML to Markdown using html-to-markdown library
	mdContent, err := markdown.ConvertString(content)
	if err != nil {
		return fmt.Errorf("convert html to markdown: %v", err)
	}

	_, err = file.WriteString(mdContent)
	if err != nil {
		return err
	}

	return nil
}

// preprocessContent fixes common formatting issues in LeetCode content
func preprocessContent(content string) string {
	// Replace &nbsp; with regular space
	content = strings.ReplaceAll(content, "&nbsp;", " ")

	// Replace multiple spaces with single space
	for strings.Contains(content, "  ") {
		content = strings.ReplaceAll(content, "  ", " ")
	}

	// Fix special HTML entities that might slip through
	content = strings.ReplaceAll(content, "&lt;", "<")
	content = strings.ReplaceAll(content, "&gt;", ">")
	content = strings.ReplaceAll(content, "&amp;", "&")
	content = strings.ReplaceAll(content, "&quot;", "\"")
	content = strings.ReplaceAll(content, "&#39;", "'")

	return content
}

// GenerateCategoryREADME generates README.md for a category directory
func GenerateCategoryREADME(basePath string, category string, problems map[string][]ProblemInfo) error {
	readmePath := filepath.Join(basePath, category, "README.md")
	file, err := os.OpenFile(readmePath, os.O_CREATE|os.O_SYNC|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write category title
	_, err = file.WriteString(fmt.Sprintf("# %s Problems\n\n", strings.ToTitle(category)))
	if err != nil {
		return err
	}

	// Write table header
	_, err = file.WriteString("| Number | Title | Solution | Test |\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("|--------|-------|----------|------|\n")
	if err != nil {
		return err
	}

	// Write problem rows
	for _, prob := range problems[category] {
		_, err = file.WriteString(fmt.Sprintf("| %s | [%s](%s) | [Go](%s) | [Test](%s) |\n",
			prob.Number,
			prob.Title,
			prob.ReadmeLink,
			prob.GoLink,
			prob.TestLink))
		if err != nil {
			return err
		}
	}

	return nil
}

// ProblemInfo holds problem metadata for README generation
type ProblemInfo struct {
	Number     string
	Title      string
	ReadmeLink string
	GoLink     string
	TestLink   string
}
