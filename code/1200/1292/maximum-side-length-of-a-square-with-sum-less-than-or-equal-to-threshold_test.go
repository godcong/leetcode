package _1292

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSideLength(t *testing.T) {
	type args struct {
		mat       [][]int
		threshold int
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
			if got := maxSideLength(tt.args.mat, tt.args.threshold); got != tt.want {
				t.Errorf("maxSideLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
