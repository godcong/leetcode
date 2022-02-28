package _1601

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumRequests(t *testing.T) {
	type args struct {
		n        int
		requests [][]int
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
			if got := maximumRequests(tt.args.n, tt.args.requests); got != tt.want {
				t.Errorf("maximumRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
