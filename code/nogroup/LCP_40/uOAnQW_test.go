package _LCP_40

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumScore(t *testing.T) {
	type args struct {
		cards []int
		cnt   int
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
			if got := maximumScore(tt.args.cards, tt.args.cnt); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
