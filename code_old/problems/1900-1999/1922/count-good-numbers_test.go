package _1922

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countGoodNumbers(t *testing.T) {
	type args struct {
		n int64
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
			if got := countGoodNumbers(tt.args.n); got != tt.want {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
