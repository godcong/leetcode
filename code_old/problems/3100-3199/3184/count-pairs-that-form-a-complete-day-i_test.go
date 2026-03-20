package _3184

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countCompleteDayPairs(t *testing.T) {
	type args struct {
		hours []int
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
			if got := countCompleteDayPairs(tt.args.hours); got != tt.want {
				t.Errorf("countCompleteDayPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
