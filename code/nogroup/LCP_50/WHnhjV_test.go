package _LCP_50

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_giveGem(t *testing.T) {
	type args struct {
		gem        []int
		operations [][]int
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
			if got := giveGem(tt.args.gem, tt.args.operations); got != tt.want {
				t.Errorf("giveGem() = %v, want %v", got, tt.want)
			}
		})
	}
}
