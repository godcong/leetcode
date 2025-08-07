package _3363

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxCollectedFruits(t *testing.T) {
	type args struct {
		fruits [][]int
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
			if got := maxCollectedFruits(tt.args.fruits); got != tt.want {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
