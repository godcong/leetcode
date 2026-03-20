package _1611

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumOneBitOperations(t *testing.T) {
	type args struct {
		n int
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
			if got := minimumOneBitOperations(tt.args.n); got != tt.want {
				t.Errorf("minimumOneBitOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
