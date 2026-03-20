package _1374

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_generateTheString(t *testing.T) {
	type args struct {
		n int
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
			if got := generateTheString(tt.args.n); got != tt.want {
				t.Errorf("generateTheString() = %v, want %v", got, tt.want)
			}
		})
	}
}
