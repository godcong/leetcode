package _0640

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_solveEquation(t *testing.T) {
	type args struct {
		equation string
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
			if got := solveEquation(tt.args.equation); got != tt.want {
				t.Errorf("solveEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
