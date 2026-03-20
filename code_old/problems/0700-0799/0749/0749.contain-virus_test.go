package _0749

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_containVirus(t *testing.T) {
	type args struct {
		isInfected [][]int
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
			if got := containVirus(tt.args.isInfected); got != tt.want {
				t.Errorf("containVirus() = %v, want %v", got, tt.want)
			}
		})
	}
}
