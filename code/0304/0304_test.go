package _304

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		matrix [][]int
		args   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				matrix: strToIntArrArray("[\n  [3, 0, 1, 4, 2],\n  [5, 6, 3, 2, 1],\n  [1, 2, 0, 1, 5],\n  [4, 1, 0, 1, 7],\n  [1, 0, 3, 0, 5]\n]"),
				args:   []int{2, 1, 4, 3},
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				matrix: strToIntArrArray("[\n  [3, 0, 1, 4, 2],\n  [5, 6, 3, 2, 1],\n  [1, 2, 0, 1, 5],\n  [4, 1, 0, 1, 7],\n  [1, 0, 3, 0, 5]\n]"),
				args:   []int{1, 1, 2, 2},
			},
			want: 11,
		},
		{
			name: "",
			args: args{
				matrix: strToIntArrArray("[\n  [3, 0, 1, 4, 2],\n  [5, 6, 3, 2, 1],\n  [1, 2, 0, 1, 5],\n  [4, 1, 0, 1, 7],\n  [1, 0, 3, 0, 5]\n]"),
				args:   []int{1, 2, 2, 4},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := Constructor(tt.args.matrix)
			got := sr.SumRegion(tt.args.args[0], tt.args.args[1], tt.args.args[2], tt.args.args[3])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
