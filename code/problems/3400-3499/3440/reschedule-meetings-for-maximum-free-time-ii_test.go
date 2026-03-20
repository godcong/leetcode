package _3440

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		startTime []int
		endTime   []int
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
			if got := maxFreeTime(tt.args.eventTime, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
