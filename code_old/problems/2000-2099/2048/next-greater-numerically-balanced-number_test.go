package _2048

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_nextBeautifulNumber(t *testing.T) {
	type args struct {
		n int
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
			if got := nextBeautifulNumber(tt.args.n); got != tt.want {
				t.Errorf("nextBeautifulNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
