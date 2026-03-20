package _2834

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumPossibleSum(t *testing.T) {
	type args struct {
		n      int
		target int
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
			if got := minimumPossibleSum(tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minimumPossibleSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
