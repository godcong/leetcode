package _2399

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkDistances(t *testing.T) {
	type args struct {
		s        string
		distance []int
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
			if got := checkDistances(tt.args.s, tt.args.distance); got != tt.want {
				t.Errorf("checkDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
