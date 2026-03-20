package _2299

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_strongPasswordCheckerII(t *testing.T) {
	type args struct {
		password string
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
			if got := strongPasswordCheckerII(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordCheckerII() = %v, want %v", got, tt.want)
			}
		})
	}
}
