package _1305

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getAllElements(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllElements(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
