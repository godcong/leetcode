package _3100

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxBottlesDrunk(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
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
			if got := maxBottlesDrunk(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("maxBottlesDrunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
