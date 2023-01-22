package _1815

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxHappyGroups(t *testing.T) {
	type args struct {
		batchSize int
		groups    []int
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
			if got := maxHappyGroups(tt.args.batchSize, tt.args.groups); got != tt.want {
				t.Errorf("maxHappyGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
