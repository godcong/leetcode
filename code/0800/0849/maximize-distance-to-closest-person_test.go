package _0849

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
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
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
