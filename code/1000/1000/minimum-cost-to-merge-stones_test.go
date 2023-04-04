package _1000

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mergeStones(t *testing.T) {
	type args struct {
		stones []int
		k      int
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
			if got := mergeStones(tt.args.stones, tt.args.k); got != tt.want {
				t.Errorf("mergeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
