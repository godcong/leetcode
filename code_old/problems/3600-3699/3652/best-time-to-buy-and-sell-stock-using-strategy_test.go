package _3652

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices   []int
		strategy []int
		k        int
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
			if got := maxProfit(tt.args.prices, tt.args.strategy, tt.args.k); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
