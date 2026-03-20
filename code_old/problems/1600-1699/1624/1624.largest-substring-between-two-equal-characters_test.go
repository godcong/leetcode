package _1624

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxLengthBetweenEqualCharacters(t *testing.T) {
	type args struct {
		s string
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
			if got := maxLengthBetweenEqualCharacters(tt.args.s); got != tt.want {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
