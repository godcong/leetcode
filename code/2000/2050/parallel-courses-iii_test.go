package _2050

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
		time      []int
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
			if got := minimumTime(tt.args.n, tt.args.relations, tt.args.time); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
