package _1079

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
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
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
