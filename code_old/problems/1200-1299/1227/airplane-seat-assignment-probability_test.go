package _1227

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_nthPersonGetsNthSeat(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPersonGetsNthSeat(tt.args.n); got != tt.want {
				t.Errorf("nthPersonGetsNthSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
