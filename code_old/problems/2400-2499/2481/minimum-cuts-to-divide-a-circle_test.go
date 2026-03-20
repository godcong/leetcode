package _2481

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfCuts(t *testing.T) {
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
			if got := numberOfCuts(tt.args.n); got != tt.want {
				t.Errorf("numberOfCuts() = %v, want %v", got, tt.want)
			}
		})
	}
}
