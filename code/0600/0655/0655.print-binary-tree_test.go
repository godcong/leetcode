package _0655

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_printTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
