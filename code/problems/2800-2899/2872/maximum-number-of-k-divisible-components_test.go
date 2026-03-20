package _2872

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxKDivisibleComponents(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		values []int
		k      int
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
			if got := maxKDivisibleComponents(tt.args.n, tt.args.edges, tt.args.values, tt.args.k); got != tt.want {
				t.Errorf("maxKDivisibleComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
