package _2920

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumPoints(t *testing.T) {
	type args struct {
		edges [][]int
		coins []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPoints(tt.args.edges, tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("maximumPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
