package _0864

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_shortestPathAllKeys(t *testing.T) {
	type args struct {
		grid []string
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
			if got := shortestPathAllKeys(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathAllKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
