package _2306

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distinctNames(t *testing.T) {
	type args struct {
		ideas []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctNames(tt.args.ideas); got != tt.want {
				t.Errorf("distinctNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
