package _2287

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rearrangeCharacters(t *testing.T) {
	type args struct {
		s      string
		target string
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
			if got := rearrangeCharacters(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("rearrangeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
