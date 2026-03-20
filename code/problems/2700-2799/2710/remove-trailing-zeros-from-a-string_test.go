package _2710

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_removeTrailingZeros(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeTrailingZeros(tt.args.num); got != tt.want {
				t.Errorf("removeTrailingZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
