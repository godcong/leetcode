package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
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
				matrix: leetcode.strToIntArrArray("[\n  [1,2,3],\n  [4,5,6],\n  [7,8,9]\n]"),
			},
			want: leetcode.strToIntArrArray("[\n  [7,4,1],\n  [8,5,2],\n  [9,6,3]\n]"),
		},
		{
			name: "",
			args: args{
				matrix: leetcode.strToIntArrArray("[\n  [ 5, 1, 9,11],\n  [ 2, 4, 8,10],\n  [13, 3, 6, 7],\n  [15,14,12,16]\n] "),
			},
			want: leetcode.strToIntArrArray("[\n  [15,13, 2, 5],\n  [14, 3, 4, 1],\n  [12, 6, 8, 9],\n  [16, 7,10,11]\n]"),
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
