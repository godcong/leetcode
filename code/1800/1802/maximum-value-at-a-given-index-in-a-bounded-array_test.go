package _1802

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxValue(t *testing.T) {
	type args struct {
		n      int
		index  int
		maxSum int
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
			if got := maxValue(tt.args.n, tt.args.index, tt.args.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
