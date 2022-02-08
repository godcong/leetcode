package old

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
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
				R:  1,
				C:  2,
				r0: 0,
				c0: 0,
			},
			want: strToIntArrArray("[[0,0],[0,1]]"),
		},
		{
			name: "",
			args: args{
				R:  2,
				C:  2,
				r0: 0,
				c0: 1,
			},
			want: strToIntArrArray("[[0,1],[0,0],[1,1],[1,0]]"),
		},
		{
			name: "",
			args: args{
				R:  2,
				C:  3,
				r0: 1,
				c0: 2,
			},
			want: strToIntArrArray("[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
