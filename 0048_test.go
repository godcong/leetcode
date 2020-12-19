package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
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
				matrix: strToIntArrArray("[\n  [1,2,3],\n  [4,5,6],\n  [7,8,9]\n]"),
			},
			want: strToIntArrArray("[\n  [7,4,1],\n  [8,5,2],\n  [9,6,3]\n]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
