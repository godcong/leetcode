package _1764

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canChoose(t *testing.T) {
	type args struct {
		groups [][]int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChoose(tt.args.groups, tt.args.nums); got != tt.want {
				t.Errorf("canChoose() = %v, want %v", got, tt.want)
			}
		})
	}
}
