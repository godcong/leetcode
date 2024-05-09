package _2105

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumRefill(t *testing.T) {
	type args struct {
		plants    []int
		capacityA int
		capacityB int
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
			if got := minimumRefill(tt.args.plants, tt.args.capacityA, tt.args.capacityB); got != tt.want {
				t.Errorf("minimumRefill() = %v, want %v", got, tt.want)
			}
		})
	}
}
