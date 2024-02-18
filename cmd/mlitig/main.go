package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	today := time.Now()
	wd, _ := os.Getwd()
	fmt.Println("Today is", today.Format("2006/01/02"))
	switch len(args) {
	case 0:
		exec.Command("gen.bat")
	case 1:
		fmt.Println("Start day is", args[0])
		start, err := time.Parse("2006/01/02", args[0])
		if err != nil {
			fmt.Println("Invalid date format", err, "date", args[0])
			return
		}
		for start.Before(today) {
			fmt.Println("Start:", start.Format("2006/01/02"))
			cmd := exec.Command(filepath.Join(wd, "gen.bat"), start.Format("2006/01/02"))
			op, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Output:", string(op))
			start = start.AddDate(0, 0, 1)
		}
	}
}
