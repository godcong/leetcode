package _2056

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countCombinations(t *testing.T) {
	type args struct {
		pieces    []string
		positions [][]int
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
			if got := countCombinations(tt.args.pieces, tt.args.positions); got != tt.want {
				t.Errorf("countCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
