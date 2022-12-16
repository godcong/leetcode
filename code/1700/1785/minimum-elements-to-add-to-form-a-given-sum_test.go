package _1785

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minElements(t *testing.T) {
	type args struct {
		nums  []int
		limit int
		goal  int
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
			if got := minElements(tt.args.nums, tt.args.limit, tt.args.goal); got != tt.want {
				t.Errorf("minElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
