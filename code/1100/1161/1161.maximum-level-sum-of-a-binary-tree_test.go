package _1161

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
