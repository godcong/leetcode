package _3180

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxTotalReward(t *testing.T) {
	type args struct {
		rewardValues []int
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
			if got := maxTotalReward(tt.args.rewardValues); got != tt.want {
				t.Errorf("maxTotalReward() = %v, want %v", got, tt.want)
			}
		})
	}
}
