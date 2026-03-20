package _2039

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_networkBecomesIdle(t *testing.T) {
	type args struct {
		edges    [][]int
		patience []int
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
			if got := networkBecomesIdle(tt.args.edges, tt.args.patience); got != tt.want {
				t.Errorf("networkBecomesIdle() = %v, want %v", got, tt.want)
			}
		})
	}
}
