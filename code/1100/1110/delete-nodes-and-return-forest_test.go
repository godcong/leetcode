package _1110

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_delNodes(t *testing.T) {
	type args struct {
		root      *TreeNode
		to_delete []int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delNodes(tt.args.root, tt.args.to_delete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
