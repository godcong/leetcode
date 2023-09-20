package _LCP_06

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minCount(t *testing.T) {
	type args struct {
		coins []int
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
			if got := minCount(tt.args.coins); got != tt.want {
				t.Errorf("minCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
