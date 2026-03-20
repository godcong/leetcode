package _SwordRefers_Offer_II_115

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sequenceReconstruction(t *testing.T) {
	type args struct {
		nums      []int
		sequences [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequenceReconstruction(tt.args.nums, tt.args.sequences); got != tt.want {
				t.Errorf("sequenceReconstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
