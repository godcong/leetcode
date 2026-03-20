package _3186

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumTotalDamage(t *testing.T) {
	type args struct {
		power []int
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
			if got := maximumTotalDamage(tt.args.power); got != tt.want {
				t.Errorf("maximumTotalDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
