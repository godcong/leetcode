package _1072

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxEqualRowsAfterFlips(t *testing.T) {
	type args struct {
		matrix [][]int
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
			if got := maxEqualRowsAfterFlips(tt.args.matrix); got != tt.want {
				t.Errorf("maxEqualRowsAfterFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
