package _1857

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_largestPathValue(t *testing.T) {
	type args struct {
		colors string
		edges  [][]int
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
			if got := largestPathValue(tt.args.colors, tt.args.edges); got != tt.want {
				t.Errorf("largestPathValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
