package _1224

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxEqualFreq(t *testing.T) {
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
			if got := maxEqualFreq(tt.args.nums); got != tt.want {
				t.Errorf("maxEqualFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}
