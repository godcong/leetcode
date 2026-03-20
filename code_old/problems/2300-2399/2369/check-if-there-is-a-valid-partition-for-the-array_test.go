package _2369

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_validPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPartition(tt.args.nums); got != tt.want {
				t.Errorf("validPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
