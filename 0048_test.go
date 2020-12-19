package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				matrix: strToIntArrArray("[\n  [1,2,3],\n  [4,5,6],\n  [7,8,9]\n]"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
