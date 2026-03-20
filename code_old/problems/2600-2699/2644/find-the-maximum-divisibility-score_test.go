package _2644

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxDivScore(t *testing.T) {
	type args struct {
		nums     []int
		divisors []int
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
			if got := maxDivScore(tt.args.nums, tt.args.divisors); got != tt.want {
				t.Errorf("maxDivScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
