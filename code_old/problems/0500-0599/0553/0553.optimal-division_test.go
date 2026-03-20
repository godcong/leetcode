package _0553

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}
