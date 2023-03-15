package _1615

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximalNetworkRank(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
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
			if got := maximalNetworkRank(tt.args.n, tt.args.roads); got != tt.want {
				t.Errorf("maximalNetworkRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
