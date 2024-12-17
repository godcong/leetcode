package _3291

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minValidStrings(t *testing.T) {
	type args struct {
		words  []string
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
			if got := minValidStrings(tt.args.words, tt.args.target); got != tt.want {
				t.Errorf("minValidStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
