package query

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/godown"
)

func WriteMarkdownTo(path string, code *Code) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_SYNC|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("### %v ###", code.Data.Question.TranslatedTitle))
	if err != nil {
		return err
	}
	file.WriteString("\n")

	reader := strings.NewReader(code.Data.Question.TranslatedContent)

	err = godown.Convert(file, reader, nil)
	if err != nil {
		return err
	}
	return nil
}
