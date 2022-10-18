package _0902

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_atMostNGivenDigitSet(t *testing.T) {
	type args struct {
		digits []string
		n      int
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
			if got := atMostNGivenDigitSet(tt.args.digits, tt.args.n); got != tt.want {
				t.Errorf("atMostNGivenDigitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
