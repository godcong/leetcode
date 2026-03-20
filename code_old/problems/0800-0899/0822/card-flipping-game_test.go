package _0822

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_flipgame(t *testing.T) {
	type args struct {
		fronts []int
		backs  []int
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
			if got := flipgame(tt.args.fronts, tt.args.backs); got != tt.want {
				t.Errorf("flipgame() = %v, want %v", got, tt.want)
			}
		})
	}
}
