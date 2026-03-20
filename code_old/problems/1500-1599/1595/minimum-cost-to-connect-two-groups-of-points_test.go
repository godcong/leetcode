package _1595

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_connectTwoGroups(t *testing.T) {
	type args struct {
		cost [][]int
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
			if got := connectTwoGroups(tt.args.cost); got != tt.want {
				t.Errorf("connectTwoGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
