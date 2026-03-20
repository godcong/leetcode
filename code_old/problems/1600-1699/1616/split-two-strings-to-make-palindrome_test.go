package _1616

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkPalindromeFormation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindromeFormation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("checkPalindromeFormation() = %v, want %v", got, tt.want)
			}
		})
	}
}
