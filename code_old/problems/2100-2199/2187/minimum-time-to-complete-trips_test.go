package _2187

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
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
			if got := minimumTime(tt.args.time, tt.args.totalTrips); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
