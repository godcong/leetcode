package _2717

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_semiOrderedPermutation(t *testing.T) {
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
			if got := semiOrderedPermutation(tt.args.nums); got != tt.want {
				t.Errorf("semiOrderedPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
