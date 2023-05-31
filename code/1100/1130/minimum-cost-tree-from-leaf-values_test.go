package _1130

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mctFromLeafValues(t *testing.T) {
	type args struct {
		arr []int
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
			if got := mctFromLeafValues(tt.args.arr); got != tt.want {
				t.Errorf("mctFromLeafValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
