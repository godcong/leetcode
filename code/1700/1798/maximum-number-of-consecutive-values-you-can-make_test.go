package _1798

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getMaximumConsecutive(t *testing.T) {
	type args struct {
		coins []int
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
			if got := getMaximumConsecutive(tt.args.coins); got != tt.want {
				t.Errorf("getMaximumConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
