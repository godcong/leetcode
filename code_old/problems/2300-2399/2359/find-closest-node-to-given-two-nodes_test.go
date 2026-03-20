package _2359

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_closestMeetingNode(t *testing.T) {
	type args struct {
		edges []int
		node1 int
		node2 int
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
			if got := closestMeetingNode(tt.args.edges, tt.args.node1, tt.args.node2); got != tt.want {
				t.Errorf("closestMeetingNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
