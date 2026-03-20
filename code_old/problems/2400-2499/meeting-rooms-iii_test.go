package _2402

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mostBooked(t *testing.T) {
	type args struct {
		n        int
		meetings [][]int
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
			if got := mostBooked(tt.args.n, tt.args.meetings); got != tt.want {
				t.Errorf("mostBooked() = %v, want %v", got, tt.want)
			}
		})
	}
}
