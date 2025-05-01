package _2071

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxTaskAssign(t *testing.T) {
	type args struct {
		tasks    []int
		workers  []int
		pills    int
		strength int
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
			if got := maxTaskAssign(tt.args.tasks, tt.args.workers, tt.args.pills, tt.args.strength); got != tt.want {
				t.Errorf("maxTaskAssign() = %v, want %v", got, tt.want)
			}
		})
	}
}
