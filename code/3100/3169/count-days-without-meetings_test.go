package _3169

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countDays(t *testing.T) {
	type args struct {
		days     int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDays(tt.args.days, tt.args.meetings); got != tt.want {
				t.Errorf("countDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
