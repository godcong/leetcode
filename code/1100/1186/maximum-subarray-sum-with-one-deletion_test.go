package _1186

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumSum(t *testing.T) {
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
			if got := maximumSum(tt.args.arr); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
