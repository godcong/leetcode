package _0998

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_insertIntoMaxTree(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
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
			if got := insertIntoMaxTree(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoMaxTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
