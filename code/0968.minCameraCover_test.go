package code

import (
	"testing"
)

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: strToTreeNode("[0,0,null,0,0]"),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				root: strToTreeNode("[0,0,null,0,null,0,null,null,0]"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
