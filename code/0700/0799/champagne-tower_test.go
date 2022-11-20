package _0799

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_champagneTower(t *testing.T) {
	type args struct {
		poured      int
		query_row   int
		query_glass int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := champagneTower(tt.args.poured, tt.args.query_row, tt.args.query_glass); got != tt.want {
				t.Errorf("champagneTower() = %v, want %v", got, tt.want)
			}
		})
	}
}
