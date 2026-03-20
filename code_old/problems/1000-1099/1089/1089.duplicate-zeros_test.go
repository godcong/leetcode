package _1089

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
		})
	}
}
