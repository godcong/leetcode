package _2698

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_punishmentNumber(t *testing.T) {
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
			if got := punishmentNumber(tt.args.n); got != tt.want {
				t.Errorf("punishmentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
