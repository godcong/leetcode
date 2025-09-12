package _3227

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_doesAliceWin(t *testing.T) {
	type args struct {
		s string
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
			if got := doesAliceWin(tt.args.s); got != tt.want {
				t.Errorf("doesAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
