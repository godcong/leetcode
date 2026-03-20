package _2952

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumAddedCoins(t *testing.T) {
	type args struct {
		coins  []int
		target int
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
			if got := minimumAddedCoins(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("minimumAddedCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
