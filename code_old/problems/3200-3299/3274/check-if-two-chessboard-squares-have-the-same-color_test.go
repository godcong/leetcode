package _3274

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkTwoChessboards(t *testing.T) {
	type args struct {
		coordinate1 string
		coordinate2 string
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
			if got := checkTwoChessboards(tt.args.coordinate1, tt.args.coordinate2); got != tt.want {
				t.Errorf("checkTwoChessboards() = %v, want %v", got, tt.want)
			}
		})
	}
}
