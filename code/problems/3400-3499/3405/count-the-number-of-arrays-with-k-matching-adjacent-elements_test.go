package _3405

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countGoodArrays(t *testing.T) {
	type args struct {
		n int
		m int
		k int
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
			if got := countGoodArrays(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("countGoodArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
