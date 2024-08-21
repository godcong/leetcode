package _3007

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findMaximumNumber(t *testing.T) {
	type args struct {
		k int64
		x int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumNumber(tt.args.k, tt.args.x); got != tt.want {
				t.Errorf("findMaximumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
