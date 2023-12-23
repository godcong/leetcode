package _1962

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minStoneSum(t *testing.T) {
	type args struct {
		piles []int
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
			if got := minStoneSum(tt.args.piles, tt.args.k); got != tt.want {
				t.Errorf("minStoneSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
