package _2065

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximalPathQuality(t *testing.T) {
	type args struct {
		values  []int
		edges   [][]int
		maxTime int
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
			if got := maximalPathQuality(tt.args.values, tt.args.edges, tt.args.maxTime); got != tt.want {
				t.Errorf("maximalPathQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
