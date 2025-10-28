package _3354

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countValidSelections(t *testing.T) {
	type args struct {
		nums []int
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
			if got := countValidSelections(tt.args.nums); got != tt.want {
				t.Errorf("countValidSelections() = %v, want %v", got, tt.want)
			}
		})
	}
}
