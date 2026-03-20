package _2081

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kMirror(t *testing.T) {
	type args struct {
		k int
		n int
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
			if got := kMirror(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("kMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}
