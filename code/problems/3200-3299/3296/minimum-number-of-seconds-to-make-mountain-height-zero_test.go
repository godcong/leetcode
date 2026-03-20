package _3296

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minNumberOfSeconds(t *testing.T) {
	type args struct {
		mountainHeight int
		workerTimes    []int
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
			if got := minNumberOfSeconds(tt.args.mountainHeight, tt.args.workerTimes); got != tt.want {
				t.Errorf("minNumberOfSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
