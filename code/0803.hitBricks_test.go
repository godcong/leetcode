package code

import (
	"reflect"
	"testing"
)

func Test_hitBricks(t *testing.T) {
	type args struct {
		grid [][]int
		hits [][]int
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
				grid: strToIntArrArray("[[1,0,0,0],[1,1,1,0]]"),
				hits: strToIntArrArray("[[1,0]]"),
			},
			want: strToIntArray("[2]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hitBricks(tt.args.grid, tt.args.hits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hitBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
