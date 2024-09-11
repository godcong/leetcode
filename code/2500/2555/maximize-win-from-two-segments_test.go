package _2555

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximizeWin(t *testing.T) {
	type args struct {
		prizePositions []int
		k              int
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
			if got := maximizeWin(tt.args.prizePositions, tt.args.k); got != tt.want {
				t.Errorf("maximizeWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
