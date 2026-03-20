package _3577

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPermutations(t *testing.T) {
	type args struct {
		complexity []int
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
			if got := countPermutations(tt.args.complexity); got != tt.want {
				t.Errorf("countPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
