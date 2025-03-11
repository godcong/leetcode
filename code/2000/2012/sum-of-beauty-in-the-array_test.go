package _2012

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sumOfBeauties(t *testing.T) {
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
			if got := sumOfBeauties(tt.args.nums); got != tt.want {
				t.Errorf("sumOfBeauties() = %v, want %v", got, tt.want)
			}
		})
	}
}
