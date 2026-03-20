package _1147

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_longestDecomposition(t *testing.T) {
	type args struct {
		text string
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
			if got := longestDecomposition(tt.args.text); got != tt.want {
				t.Errorf("longestDecomposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
