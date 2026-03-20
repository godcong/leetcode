package _0865

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_subtreeWithAllDeepest(t *testing.T) {
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
			if got := subtreeWithAllDeepest(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtreeWithAllDeepest() = %v, want %v", got, tt.want)
			}
		})
	}
}
