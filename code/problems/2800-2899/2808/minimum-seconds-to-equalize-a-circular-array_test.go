package _2808

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumSeconds(t *testing.T) {
	type args struct {
		nums []int
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
			if got := minimumSeconds(tt.args.nums); got != tt.want {
				t.Errorf("minimumSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
