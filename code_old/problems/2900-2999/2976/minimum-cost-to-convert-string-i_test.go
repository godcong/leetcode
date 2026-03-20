package _2976

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
		source   string
		target   string
		original []byte
		changed  []byte
		cost     []int
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
			if got := minimumCost(tt.args.source, tt.args.target, tt.args.original, tt.args.changed, tt.args.cost); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
