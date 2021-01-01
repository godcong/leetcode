package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
