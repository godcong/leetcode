package _3001

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minMovesToCaptureTheQueen(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
		e int
		f int
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
			if got := minMovesToCaptureTheQueen(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e, tt.args.f); got != tt.want {
				t.Errorf("minMovesToCaptureTheQueen() = %v, want %v", got, tt.want)
			}
		})
	}
}
