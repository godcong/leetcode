package _3098

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sumOfPowers(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := sumOfPowers(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("sumOfPowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
