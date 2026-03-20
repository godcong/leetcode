package _0781

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
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
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
