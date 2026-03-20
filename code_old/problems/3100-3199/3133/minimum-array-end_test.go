package _3133

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minEnd(t *testing.T) {
	type args struct {
		n int
		x int
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
			if got := minEnd(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("minEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
