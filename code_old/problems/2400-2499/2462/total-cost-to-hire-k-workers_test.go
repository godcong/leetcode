package _2462

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_totalCost(t *testing.T) {
	type args struct {
		costs      []int
		k          int
		candidates int
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
			if got := totalCost(tt.args.costs, tt.args.k, tt.args.candidates); got != tt.want {
				t.Errorf("totalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
