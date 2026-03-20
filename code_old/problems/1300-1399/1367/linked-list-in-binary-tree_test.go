package _1367

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_isSubPath(t *testing.T) {
	type args struct {
		head *ListNode
		root *TreeNode
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
			if got := isSubPath(tt.args.head, tt.args.root); got != tt.want {
				t.Errorf("isSubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
