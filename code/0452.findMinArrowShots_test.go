package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
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
				points: leetcode.strToIntArrArray("[[10,16],[2,8],[1,6],[7,12]]"),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				points: leetcode.strToIntArrArray("[[1,2],[3,4],[5,6],[7,8]]"),
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				points: leetcode.strToIntArrArray("[[1,2],[2,3],[3,4],[4,5]]"),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				points: leetcode.strToIntArrArray("[[1,2]]"),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				points: leetcode.strToIntArrArray("[[2,3],[2,3]]"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
