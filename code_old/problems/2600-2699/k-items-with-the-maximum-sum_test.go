package _2600

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kItemsWithMaximumSum(t *testing.T) {
	type args struct {
		numOnes    int
		numZeros   int
		numNegOnes int
		k          int
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
			if got := kItemsWithMaximumSum(tt.args.numOnes, tt.args.numZeros, tt.args.numNegOnes, tt.args.k); got != tt.want {
				t.Errorf("kItemsWithMaximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
