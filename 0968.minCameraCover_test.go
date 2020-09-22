package leetcode

import "testing"

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
