package _2959

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfSets(t *testing.T) {
	type args struct {
		n           int
		maxDistance int
		roads       [][]int
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
			if got := numberOfSets(tt.args.n, tt.args.maxDistance, tt.args.roads); got != tt.want {
				t.Errorf("numberOfSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
