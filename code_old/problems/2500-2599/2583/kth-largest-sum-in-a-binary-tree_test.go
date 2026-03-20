package _2583

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kthLargestLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
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
			if got := kthLargestLevelSum(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargestLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
