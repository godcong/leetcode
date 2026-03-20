package _2741

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_specialPerm(t *testing.T) {
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
			if got := specialPerm(tt.args.nums); got != tt.want {
				t.Errorf("specialPerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
