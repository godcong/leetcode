package _0623

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
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
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
