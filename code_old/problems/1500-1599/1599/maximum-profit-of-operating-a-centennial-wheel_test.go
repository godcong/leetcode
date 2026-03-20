package _1599

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minOperationsMaxProfit(t *testing.T) {
	type args struct {
		customers    []int
		boardingCost int
		runningCost  int
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
			if got := minOperationsMaxProfit(tt.args.customers, tt.args.boardingCost, tt.args.runningCost); got != tt.want {
				t.Errorf("minOperationsMaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
