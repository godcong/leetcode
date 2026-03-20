package _2275

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_largestCombination(t *testing.T) {
	type args struct {
		candidates []int
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
			if got := largestCombination(tt.args.candidates); got != tt.want {
				t.Errorf("largestCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
