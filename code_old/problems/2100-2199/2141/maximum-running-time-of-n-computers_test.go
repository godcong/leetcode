package _2141

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxRunTime(t *testing.T) {
	type args struct {
		n         int
		batteries []int
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
			if got := maxRunTime(tt.args.n, tt.args.batteries); got != tt.want {
				t.Errorf("maxRunTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
