package _2274

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxConsecutive(t *testing.T) {
	type args struct {
		bottom  int
		top     int
		special []int
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
			if got := maxConsecutive(tt.args.bottom, tt.args.top, tt.args.special); got != tt.want {
				t.Errorf("maxConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
