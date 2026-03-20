package _2561

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minCost(t *testing.T) {
	type args struct {
		basket1 []int
		basket2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.basket1, tt.args.basket2); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
