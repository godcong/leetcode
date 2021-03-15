package _0054

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				matrix: strToIntArrArray("[[1,2,3],[4,5,6],[7,8,9]]"),
			},
			want: strToIntArray("[1,2,3,6,9,8,7,4,5]"),
		},
		{
			name: "",
			args: args{
				matrix: strToIntArrArray("[[1,2,3,4],[5,6,7,8],[9,10,11,12]]"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
