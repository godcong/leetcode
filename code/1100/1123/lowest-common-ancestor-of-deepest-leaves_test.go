package _1123

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_lcaDeepestLeaves(t *testing.T) {
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
			if got := lcaDeepestLeaves(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lcaDeepestLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
