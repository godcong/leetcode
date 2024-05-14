package _2244

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumRounds(t *testing.T) {
	type args struct {
		tasks []int
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
			if got := minimumRounds(tt.args.tasks); got != tt.want {
				t.Errorf("minimumRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
