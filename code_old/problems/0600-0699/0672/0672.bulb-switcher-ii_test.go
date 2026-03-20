package _0672

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_flipLights(t *testing.T) {
	type args struct {
		n       int
		presses int
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
			if got := flipLights(tt.args.n, tt.args.presses); got != tt.want {
				t.Errorf("flipLights() = %v, want %v", got, tt.want)
			}
		})
	}
}
