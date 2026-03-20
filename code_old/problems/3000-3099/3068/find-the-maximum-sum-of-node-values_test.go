package _3068

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumValueSum(t *testing.T) {
	type args struct {
		nums  []int
		k     int
		edges [][]int
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
			if got := maximumValueSum(tt.args.nums, tt.args.k, tt.args.edges); got != tt.want {
				t.Errorf("maximumValueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
