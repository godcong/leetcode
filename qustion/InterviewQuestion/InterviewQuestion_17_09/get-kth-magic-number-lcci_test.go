package _InterviewQuestion_17_09

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getKthMagicNumber(t *testing.T) {
	type args struct {
		k int
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
			if got := getKthMagicNumber(tt.args.k); got != tt.want {
				t.Errorf("getKthMagicNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
