package _0324

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
		})
	}
}
