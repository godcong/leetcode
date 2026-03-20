package _2931

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSpending(t *testing.T) {
	type args struct {
		values [][]int
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
			if got := maxSpending(tt.args.values); got != tt.want {
				t.Errorf("maxSpending() = %v, want %v", got, tt.want)
			}
		})
	}
}
