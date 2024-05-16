package _1953

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfWeeks(t *testing.T) {
	type args struct {
		milestones []int
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
			if got := numberOfWeeks(tt.args.milestones); got != tt.want {
				t.Errorf("numberOfWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}
