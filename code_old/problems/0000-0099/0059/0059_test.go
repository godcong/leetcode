package _0059

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
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
				n: 3,
			},
			want: StrToIntArrArray("[[1,2,3],[8,9,4],[7,6,5]]"),
		},
		{
			name: "",
			args: args{
				n: 1,
			},
			want: StrToIntArrArray("[[1]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
