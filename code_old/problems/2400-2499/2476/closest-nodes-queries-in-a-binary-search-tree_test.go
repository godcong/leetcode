package _2476

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_closestNodes(t *testing.T) {
	type args struct {
		root    *TreeNode
		queries []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestNodes(tt.args.root, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
