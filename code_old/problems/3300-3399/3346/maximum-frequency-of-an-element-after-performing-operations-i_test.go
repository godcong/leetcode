package _3346

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums          []int
		k             int
		numOperations int
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
			if got := maxFrequency(tt.args.nums, tt.args.k, tt.args.numOperations); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
