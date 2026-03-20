package _2532

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findCrossingTime(t *testing.T) {
	type args struct {
		n    int
		k    int
		time [][]int
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
			if got := findCrossingTime(tt.args.n, tt.args.k, tt.args.time); got != tt.want {
				t.Errorf("findCrossingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
