package _3541

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxFreqSum(t *testing.T) {
	type args struct {
		s string
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
			if got := maxFreqSum(tt.args.s); got != tt.want {
				t.Errorf("maxFreqSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
