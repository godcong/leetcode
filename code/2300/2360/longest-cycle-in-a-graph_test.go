package _2360

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_longestCycle(t *testing.T) {
	type args struct {
		edges []int
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
			if got := longestCycle(tt.args.edges); got != tt.want {
				t.Errorf("longestCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
