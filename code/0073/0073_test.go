package _0073

import (
	. "github.com/godcong/leetcode/common"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
