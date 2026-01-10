package _0712

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
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
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
