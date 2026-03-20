package _2210

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countHillValley(t *testing.T) {
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
			if got := countHillValley(tt.args.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
