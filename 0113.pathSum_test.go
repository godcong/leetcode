package leetcode

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: strToTreeNode("[5,4,8,11,null,13,4,7,2,null,null,5,1]"),
				sum:  22,
			},
			want: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
