package _2140

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mostPoints(t *testing.T) {
	type args struct {
		questions [][]int
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
			if got := mostPoints(tt.args.questions); got != tt.want {
				t.Errorf("mostPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
