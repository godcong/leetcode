package _2528

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxPower(t *testing.T) {
	type args struct {
		stations []int
		r        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.stations, tt.args.r, tt.args.k); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
