package _2379

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
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
			if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
