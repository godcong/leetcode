package _2673

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minIncrements(t *testing.T) {
	type args struct {
		n    int
		cost []int
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
			if got := minIncrements(tt.args.n, tt.args.cost); got != tt.want {
				t.Errorf("minIncrements() = %v, want %v", got, tt.want)
			}
		})
	}
}
