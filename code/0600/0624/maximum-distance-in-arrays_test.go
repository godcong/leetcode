package _0624

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
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
			if got := maxDistance(tt.args.arrays); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
