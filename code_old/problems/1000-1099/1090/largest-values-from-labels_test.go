package _1090

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_largestValsFromLabels(t *testing.T) {
	type args struct {
		values    []int
		labels    []int
		numWanted int
		useLimit  int
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
			if got := largestValsFromLabels(tt.args.values, tt.args.labels, tt.args.numWanted, tt.args.useLimit); got != tt.want {
				t.Errorf("largestValsFromLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
