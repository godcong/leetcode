package _0777

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canTransform(t *testing.T) {
	type args struct {
		start string
		end   string
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
			if got := canTransform(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("canTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
