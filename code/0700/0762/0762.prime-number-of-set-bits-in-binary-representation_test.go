package _0762

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		left  int
		right int
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
			if got := countPrimeSetBits(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
