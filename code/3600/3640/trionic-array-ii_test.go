package _3640

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSumTrionic(t *testing.T) {
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
			if got := maxSumTrionic(tt.args.nums); got != tt.want {
				t.Errorf("maxSumTrionic() = %v, want %v", got, tt.want)
			}
		})
	}
}
