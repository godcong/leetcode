package _2603

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_collectTheCoins(t *testing.T) {
	type args struct {
		coins []int
		edges [][]int
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
			if got := collectTheCoins(tt.args.coins, tt.args.edges); got != tt.want {
				t.Errorf("collectTheCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
