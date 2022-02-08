package _0073

import (
	. "github.com/godcong/leetcode/common"
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
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
				matrix: StrToIntArrArray("[[1,1,1],[1,0,1],[1,1,1]]"),
			},
			want: StrToIntArrArray("[[1,0,1],[0,0,0],[1,0,1]]"),
		},
		{
			name: "",
			args: args{
				matrix: StrToIntArrArray("[[0,1,2,0],[3,4,5,2],[1,3,1,5]]"),
			},
			want: StrToIntArrArray("[[0,0,0,0],[0,4,5,0],[0,3,1,0]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.want)
			}

		})
	}
}
