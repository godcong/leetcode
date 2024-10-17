package _3193

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfPermutations(t *testing.T) {
	type args struct {
		n            int
		requirements [][]int
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
			if got := numberOfPermutations(tt.args.n, tt.args.requirements); got != tt.want {
				t.Errorf("numberOfPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
