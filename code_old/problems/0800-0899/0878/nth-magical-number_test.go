package _0878

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_nthMagicalNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
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
			if got := nthMagicalNumber(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("nthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
