package _3151

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums); got != tt.want {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
