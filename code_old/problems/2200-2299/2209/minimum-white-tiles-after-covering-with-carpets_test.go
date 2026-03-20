package _2209

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumWhiteTiles(t *testing.T) {
	type args struct {
		floor      string
		numCarpets int
		carpetLen  int
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
			if got := minimumWhiteTiles(tt.args.floor, tt.args.numCarpets, tt.args.carpetLen); got != tt.want {
				t.Errorf("minimumWhiteTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
