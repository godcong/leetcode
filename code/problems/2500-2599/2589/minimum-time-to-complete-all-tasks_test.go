package _2589

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findMinimumTime(t *testing.T) {
	type args struct {
		tasks [][]int
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
			if got := findMinimumTime(tt.args.tasks); got != tt.want {
				t.Errorf("findMinimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
