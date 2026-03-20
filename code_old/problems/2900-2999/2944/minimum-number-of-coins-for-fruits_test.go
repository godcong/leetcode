package _2944

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumCoins(t *testing.T) {
	type args struct {
		prices []int
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
			if got := minimumCoins(tt.args.prices); got != tt.want {
				t.Errorf("minimumCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
