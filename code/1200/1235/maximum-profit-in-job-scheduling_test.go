package _1235

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_jobScheduling(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		profit    []int
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
			if got := jobScheduling(tt.args.startTime, tt.args.endTime, tt.args.profit); got != tt.want {
				t.Errorf("jobScheduling() = %v, want %v", got, tt.want)
			}
		})
	}
}
