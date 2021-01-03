package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
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
				cost: strToIntArray("[10, 15, 20]"),
			},
			want: 15,
		},
		{
			name: "",
			args: args{
				cost: strToIntArray("[1, 100, 1, 1, 1, 100, 1, 1, 100, 1]"),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
