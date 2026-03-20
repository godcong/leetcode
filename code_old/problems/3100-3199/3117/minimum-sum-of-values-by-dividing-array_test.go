package _3117

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumValueSum(t *testing.T) {
	type args struct {
		nums      []int
		andValues []int
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
			if got := minimumValueSum(tt.args.nums, tt.args.andValues); got != tt.want {
				t.Errorf("minimumValueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
