package _0736

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_evaluate(t *testing.T) {
	type args struct {
		expression string
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
			if got := evaluate(tt.args.expression); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
