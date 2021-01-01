package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
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
			want: []float64{
				3, 14.5, 11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
