package _1154

import (
	"time"
)

const yearDate = "2006-01-02"

func dayOfYear(date string) int {
	t, _ := time.Parse(yearDate, date)
	return t.YearDay()
}
