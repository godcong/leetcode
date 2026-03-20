package _3024

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
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
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
