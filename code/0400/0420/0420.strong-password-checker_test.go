package _0420

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_strongPasswordChecker(t *testing.T) {
	type args struct {
		password string
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
			if got := strongPasswordChecker(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
