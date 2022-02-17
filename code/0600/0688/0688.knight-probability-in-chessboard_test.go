package _0688

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_knightProbability(t *testing.T) {
	type args struct {
		n      int
		k      int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightProbability(tt.args.n, tt.args.k, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
