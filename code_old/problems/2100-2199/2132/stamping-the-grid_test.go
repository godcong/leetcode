package _2132

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_possibleToStamp(t *testing.T) {
	type args struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleToStamp(tt.args.grid, tt.args.stampHeight, tt.args.stampWidth); got != tt.want {
				t.Errorf("possibleToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
