package _0814

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
