package _2218

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxValueOfCoins(t *testing.T) {
	type args struct {
		piles [][]int
		k     int
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
			if got := maxValueOfCoins(tt.args.piles, tt.args.k); got != tt.want {
				t.Errorf("maxValueOfCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
