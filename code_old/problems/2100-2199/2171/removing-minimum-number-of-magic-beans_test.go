package _2171

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumRemoval(t *testing.T) {
	type args struct {
		beans []int
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
			if got := minimumRemoval(tt.args.beans); got != tt.want {
				t.Errorf("minimumRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
