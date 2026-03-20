package _2364

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countBadPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBadPairs(tt.args.nums); got != tt.want {
				t.Errorf("countBadPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
