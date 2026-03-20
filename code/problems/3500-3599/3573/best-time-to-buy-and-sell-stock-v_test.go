package _3573

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumProfit(t *testing.T) {
	type args struct {
		prices []int
		k      int
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
			if got := maximumProfit(tt.args.prices, tt.args.k); got != tt.want {
				t.Errorf("maximumProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
