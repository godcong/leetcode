package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				root: &leetcode.TreeNode{
					Val: 3,
					Left: &leetcode.TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val:   20,
						Left:  &leetcode.TreeNode{Val: 15},
						Right: &leetcode.TreeNode{Val: 7},
					},
				},
			},
			want: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
