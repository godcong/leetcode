package _3343

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countBalancedPermutations(t *testing.T) {
	type args struct {
		num string
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
			if got := countBalancedPermutations(tt.args.num); got != tt.want {
				t.Errorf("countBalancedPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
