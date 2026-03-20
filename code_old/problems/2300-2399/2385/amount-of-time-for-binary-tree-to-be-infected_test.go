package _2385

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_amountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
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
			if got := amountOfTime(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
