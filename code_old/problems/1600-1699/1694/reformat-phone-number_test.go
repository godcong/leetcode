package _1694

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reformatNumber(t *testing.T) {
	type args struct {
		number string
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
			if got := reformatNumber(tt.args.number); got != tt.want {
				t.Errorf("reformatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
