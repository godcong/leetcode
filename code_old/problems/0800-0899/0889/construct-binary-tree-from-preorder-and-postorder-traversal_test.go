package _0889

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_constructFromPrePost(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
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
			if got := constructFromPrePost(tt.args.preorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
