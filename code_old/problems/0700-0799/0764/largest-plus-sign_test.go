package _0764

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_orderOfLargestPlusSign(t *testing.T) {
	type args struct {
		n     int
		mines [][]int
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
			if got := orderOfLargestPlusSign(tt.args.n, tt.args.mines); got != tt.want {
				t.Errorf("orderOfLargestPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
