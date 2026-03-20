package _1263

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minPushBox(t *testing.T) {
	type args struct {
		grid [][]byte
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
			if got := minPushBox(tt.args.grid); got != tt.want {
				t.Errorf("minPushBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
