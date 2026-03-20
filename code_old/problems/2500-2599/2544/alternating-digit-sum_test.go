package _2544

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_alternateDigitSum(t *testing.T) {
	type args struct {
		n int
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
			if got := alternateDigitSum(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
