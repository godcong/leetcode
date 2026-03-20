package _2347

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_bestHand(t *testing.T) {
	type args struct {
		ranks []int
		suits []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestHand(tt.args.ranks, tt.args.suits); got != tt.want {
				t.Errorf("bestHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
