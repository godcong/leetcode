package code

import (
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: strToIntArrArray("[[0,1,0,0],\n [1,1,1,0],\n [0,1,0,0],\n [1,1,0,0]]"),
			},
			want: 16,
		},
		{
			name: "",
			args: args{
				grid: strToIntArrArray("[[1,0]]"),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
