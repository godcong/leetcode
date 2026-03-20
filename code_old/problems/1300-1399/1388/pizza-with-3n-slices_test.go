package _1388

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSizeSlices(t *testing.T) {
	type args struct {
		slices []int
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
			if got := maxSizeSlices(tt.args.slices); got != tt.want {
				t.Errorf("maxSizeSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
