package _1781

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_beautySum(t *testing.T) {
	type args struct {
		s string
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
			if got := beautySum(tt.args.s); got != tt.want {
				t.Errorf("beautySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
