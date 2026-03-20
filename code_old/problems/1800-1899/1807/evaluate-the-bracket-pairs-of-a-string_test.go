package _1807

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
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
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
