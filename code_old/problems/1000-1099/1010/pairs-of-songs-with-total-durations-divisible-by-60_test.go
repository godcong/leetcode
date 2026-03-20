package _1010

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
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
			if got := numPairsDivisibleBy60(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
