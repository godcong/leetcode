package _0952

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_largestComponentSize(t *testing.T) {
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
			if got := largestComponentSize(tt.args.nums); got != tt.want {
				t.Errorf("largestComponentSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
